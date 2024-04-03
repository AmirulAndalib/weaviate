//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package offload

import (
	"context"
	"errors"
	"fmt"
	"time"

	enterrors "github.com/weaviate/weaviate/entities/errors"

	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/entities/offload"
	"github.com/weaviate/weaviate/usecases/config"
)

type backupper struct {
	node     string
	logger   logrus.FieldLogger
	sourcer  Sourcer
	backends OffloadBackendProvider
	// shardCoordinationChan is sync and coordinate operations
	shardSyncChan
}

func newBackupper(node string, logger logrus.FieldLogger, sourcer Sourcer, backends OffloadBackendProvider,
) *backupper {
	return &backupper{
		node:          node,
		logger:        logger,
		sourcer:       sourcer,
		backends:      backends,
		shardSyncChan: shardSyncChan{coordChan: make(chan interface{}, 5)},
	}
}

// // Backup is called by the User
// func (b *backupper) Backup(ctx context.Context,
// 	store nodeStore, id string, classes []string,
// ) (*backup.CreateMeta, error) {
// 	// make sure there is no active backup
// 	req := Request{
// 		Method:  OpCreate,
// 		ID:      id,
// 		Classes: classes,
// 	}
// 	if _, err := b.backup(ctx, store, &req); err != nil {
// 		return nil, backup.NewErrUnprocessable(err)
// 	}

// 	return &backup.CreateMeta{
// 		Path:   store.HomeDir(),
// 		Status: backup.Started,
// 	}, nil
// }

// // Status returns status of a backup
// // If the backup is still active the status is immediately returned
// // If not it fetches the metadata file to get the status
// func (b *backupper) Status(ctx context.Context, backend, bakID string,
// ) (*models.BackupCreateStatusResponse, error) {
// 	st, err := b.OnStatus(ctx, &StatusRequest{OpCreate, bakID, backend})
// 	if err != nil {
// 		if errors.Is(err, errMetaNotFound) {
// 			err = backup.NewErrNotFound(err)
// 		} else {
// 			err = backup.NewErrUnprocessable(err)
// 		}
// 		return nil, err
// 	}
// 	// check if backup is still active
// 	status := string(st.Status)
// 	return &models.BackupCreateStatusResponse{
// 		ID:      bakID,
// 		Path:    st.Path,
// 		Status:  &status,
// 		Backend: backend,
// 	}, nil
// }

func (b *backupper) OnStatus(ctx context.Context, req *StatusRequest) (reqStat, error) {
	// check if backup is still active
	st := b.lastOp.get()
	if st.ID == req.ID {
		return st, nil
	}

	// The backup might have been already created.
	store, err := nodeBackend(b.node, b.backends, req.Backend, req.ID)
	if err != nil {
		return reqStat{}, fmt.Errorf("no backup provider %q, did you enable the right module?", req.Backend)
	}

	meta, err := store.Meta(ctx)
	if err != nil {
		path := fmt.Sprintf("%s/%s", req.ID, OffloadFile)
		return reqStat{}, fmt.Errorf("cannot get status while backing up: %w: %q: %v", errMetaNotFound, path, err)
	}
	if err != nil || meta.Error != "" {
		return reqStat{}, errors.New(meta.Error)
	}

	return reqStat{
		StartTime: meta.StartedAt,
		ID:        req.ID,
		Path:      store.HomeDir(),
		Status:    offload.Status(meta.Status),
	}, nil
}

// backup checks if the node is ready to back up (can commit phase)
//
// Moreover it starts a goroutine in the background which waits for the
// next instruction from the coordinator (second phase).
// It will start the backup as soon as it receives an ack, or abort otherwise
func (b *backupper) backup(ctx context.Context,
	store nodeStore, req *Request,
) (CanCommitResponse, error) {
	expiration := req.Duration
	if expiration > _TimeoutShardCommit {
		expiration = _TimeoutShardCommit
	}
	ret := CanCommitResponse{
		Method:  OpOffload,
		ID:      req.ID,
		Timeout: expiration,
	}
	// make sure there is no active backup
	if prevID := b.lastOp.renew(req.ID, store.HomeDir()); prevID != "" {
		return ret, fmt.Errorf("backup %s already in progress", prevID)
	}
	b.waitingForCoordinatorToCommit.Store(true) // is set to false by wait()
	// waits for ack from coordinator in order to processed with the backup
	f := func() {
		defer b.lastOp.reset()
		if err := b.waitForCoordinator(expiration, req.ID); err != nil {
			b.logger.WithField("action", "create_backup").
				Error(err)
			b.lastAsyncError = err
			return

		}
		uploader := newUploader(b.sourcer, store, b.lastOp.set, b.logger).
			withCompression(newZipConfig(req.Compression))

		result := offload.OffloadNodeDescriptor{
			StartedAt:     time.Now().UTC(),
			ID:            req.ID,
			Class:         req.Class,
			Tenant:        req.Tenant,
			Version:       Version,
			ServerVersion: config.ServerVersion,
		}

		// the coordinator might want to abort the backup
		done := make(chan struct{})
		ctx := b.withCancellation(context.Background(), req.ID, done)
		defer close(done)

		logFields := logrus.Fields{"action": "create_backup", "backup_id": req.ID}
		if err := uploader.all(ctx, req.Class, req.Tenant, &result); err != nil {
			b.logger.WithFields(logFields).Error(err)
			b.lastAsyncError = err

		} else {
			b.logger.WithFields(logFields).Info("backup completed successfully")
		}
		result.CompletedAt = time.Now().UTC()
	}
	enterrors.GoWrapper(f, b.logger)

	return ret, nil
}
//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationReplicateDetailsReplicaStatusError Represents an error encountered during a replication operation, including its timestamp and a human-readable message.
//
// swagger:model ReplicationReplicateDetailsReplicaStatusError
type ReplicationReplicateDetailsReplicaStatusError struct {

	// A human-readable message describing the error.
	Message string `json:"message,omitempty"`

	// The unix timestamp in ms when the error occurred. This is an approximate time and so should not be used for precise timing.
	WhenErroredUnixMs int64 `json:"whenErroredUnixMs,omitempty"`
}

// Validate validates this replication replicate details replica status error
func (m *ReplicationReplicateDetailsReplicaStatusError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replication replicate details replica status error based on context it is used
func (m *ReplicationReplicateDetailsReplicaStatusError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationReplicateDetailsReplicaStatusError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationReplicateDetailsReplicaStatusError) UnmarshalBinary(b []byte) error {
	var res ReplicationReplicateDetailsReplicaStatusError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

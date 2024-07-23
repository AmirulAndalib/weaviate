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

// Code generated by mockery v2.43.2. DO NOT EDIT.

package scaler

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Mockclient is an autogenerated mock type for the client type
type Mockclient struct {
	mock.Mock
}

type Mockclient_Expecter struct {
	mock *mock.Mock
}

func (_m *Mockclient) EXPECT() *Mockclient_Expecter {
	return &Mockclient_Expecter{mock: &_m.Mock}
}

// CreateShard provides a mock function with given fields: ctx, hostName, indexName, shardName
func (_m *Mockclient) CreateShard(ctx context.Context, hostName string, indexName string, shardName string) error {
	ret := _m.Called(ctx, hostName, indexName, shardName)

	if len(ret) == 0 {
		panic("no return value specified for CreateShard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, hostName, indexName, shardName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mockclient_CreateShard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateShard'
type Mockclient_CreateShard_Call struct {
	*mock.Call
}

// CreateShard is a helper method to define mock.On call
//   - ctx context.Context
//   - hostName string
//   - indexName string
//   - shardName string
func (_e *Mockclient_Expecter) CreateShard(ctx interface{}, hostName interface{}, indexName interface{}, shardName interface{}) *Mockclient_CreateShard_Call {
	return &Mockclient_CreateShard_Call{Call: _e.mock.On("CreateShard", ctx, hostName, indexName, shardName)}
}

func (_c *Mockclient_CreateShard_Call) Run(run func(ctx context.Context, hostName string, indexName string, shardName string)) *Mockclient_CreateShard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *Mockclient_CreateShard_Call) Return(_a0 error) *Mockclient_CreateShard_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockclient_CreateShard_Call) RunAndReturn(run func(context.Context, string, string, string) error) *Mockclient_CreateShard_Call {
	_c.Call.Return(run)
	return _c
}

// IncreaseReplicationFactor provides a mock function with given fields: ctx, host, class, dist
func (_m *Mockclient) IncreaseReplicationFactor(ctx context.Context, host string, class string, dist ShardDist) error {
	ret := _m.Called(ctx, host, class, dist)

	if len(ret) == 0 {
		panic("no return value specified for IncreaseReplicationFactor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ShardDist) error); ok {
		r0 = rf(ctx, host, class, dist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mockclient_IncreaseReplicationFactor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncreaseReplicationFactor'
type Mockclient_IncreaseReplicationFactor_Call struct {
	*mock.Call
}

// IncreaseReplicationFactor is a helper method to define mock.On call
//   - ctx context.Context
//   - host string
//   - class string
//   - dist ShardDist
func (_e *Mockclient_Expecter) IncreaseReplicationFactor(ctx interface{}, host interface{}, class interface{}, dist interface{}) *Mockclient_IncreaseReplicationFactor_Call {
	return &Mockclient_IncreaseReplicationFactor_Call{Call: _e.mock.On("IncreaseReplicationFactor", ctx, host, class, dist)}
}

func (_c *Mockclient_IncreaseReplicationFactor_Call) Run(run func(ctx context.Context, host string, class string, dist ShardDist)) *Mockclient_IncreaseReplicationFactor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(ShardDist))
	})
	return _c
}

func (_c *Mockclient_IncreaseReplicationFactor_Call) Return(_a0 error) *Mockclient_IncreaseReplicationFactor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockclient_IncreaseReplicationFactor_Call) RunAndReturn(run func(context.Context, string, string, ShardDist) error) *Mockclient_IncreaseReplicationFactor_Call {
	_c.Call.Return(run)
	return _c
}

// PutFile provides a mock function with given fields: ctx, hostName, indexName, shardName, fileName, payload
func (_m *Mockclient) PutFile(ctx context.Context, hostName string, indexName string, shardName string, fileName string, payload io.ReadSeekCloser) error {
	ret := _m.Called(ctx, hostName, indexName, shardName, fileName, payload)

	if len(ret) == 0 {
		panic("no return value specified for PutFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, io.ReadSeekCloser) error); ok {
		r0 = rf(ctx, hostName, indexName, shardName, fileName, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mockclient_PutFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutFile'
type Mockclient_PutFile_Call struct {
	*mock.Call
}

// PutFile is a helper method to define mock.On call
//   - ctx context.Context
//   - hostName string
//   - indexName string
//   - shardName string
//   - fileName string
//   - payload io.ReadSeekCloser
func (_e *Mockclient_Expecter) PutFile(ctx interface{}, hostName interface{}, indexName interface{}, shardName interface{}, fileName interface{}, payload interface{}) *Mockclient_PutFile_Call {
	return &Mockclient_PutFile_Call{Call: _e.mock.On("PutFile", ctx, hostName, indexName, shardName, fileName, payload)}
}

func (_c *Mockclient_PutFile_Call) Run(run func(ctx context.Context, hostName string, indexName string, shardName string, fileName string, payload io.ReadSeekCloser)) *Mockclient_PutFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(io.ReadSeekCloser))
	})
	return _c
}

func (_c *Mockclient_PutFile_Call) Return(_a0 error) *Mockclient_PutFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockclient_PutFile_Call) RunAndReturn(run func(context.Context, string, string, string, string, io.ReadSeekCloser) error) *Mockclient_PutFile_Call {
	_c.Call.Return(run)
	return _c
}

// ReInitShard provides a mock function with given fields: ctx, hostName, indexName, shardName
func (_m *Mockclient) ReInitShard(ctx context.Context, hostName string, indexName string, shardName string) error {
	ret := _m.Called(ctx, hostName, indexName, shardName)

	if len(ret) == 0 {
		panic("no return value specified for ReInitShard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, hostName, indexName, shardName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mockclient_ReInitShard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReInitShard'
type Mockclient_ReInitShard_Call struct {
	*mock.Call
}

// ReInitShard is a helper method to define mock.On call
//   - ctx context.Context
//   - hostName string
//   - indexName string
//   - shardName string
func (_e *Mockclient_Expecter) ReInitShard(ctx interface{}, hostName interface{}, indexName interface{}, shardName interface{}) *Mockclient_ReInitShard_Call {
	return &Mockclient_ReInitShard_Call{Call: _e.mock.On("ReInitShard", ctx, hostName, indexName, shardName)}
}

func (_c *Mockclient_ReInitShard_Call) Run(run func(ctx context.Context, hostName string, indexName string, shardName string)) *Mockclient_ReInitShard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *Mockclient_ReInitShard_Call) Return(_a0 error) *Mockclient_ReInitShard_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockclient_ReInitShard_Call) RunAndReturn(run func(context.Context, string, string, string) error) *Mockclient_ReInitShard_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockclient creates a new instance of Mockclient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockclient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mockclient {
	mock := &Mockclient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

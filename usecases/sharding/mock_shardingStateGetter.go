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

package sharding

import mock "github.com/stretchr/testify/mock"

// MockshardingStateGetter is an autogenerated mock type for the shardingStateGetter type
type MockshardingStateGetter struct {
	mock.Mock
}

type MockshardingStateGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockshardingStateGetter) EXPECT() *MockshardingStateGetter_Expecter {
	return &MockshardingStateGetter_Expecter{mock: &_m.Mock}
}

// ShardOwner provides a mock function with given fields: class, shard
func (_m *MockshardingStateGetter) ShardOwner(class string, shard string) (string, error) {
	ret := _m.Called(class, shard)

	if len(ret) == 0 {
		panic("no return value specified for ShardOwner")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(class, shard)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(class, shard)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(class, shard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockshardingStateGetter_ShardOwner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShardOwner'
type MockshardingStateGetter_ShardOwner_Call struct {
	*mock.Call
}

// ShardOwner is a helper method to define mock.On call
//   - class string
//   - shard string
func (_e *MockshardingStateGetter_Expecter) ShardOwner(class interface{}, shard interface{}) *MockshardingStateGetter_ShardOwner_Call {
	return &MockshardingStateGetter_ShardOwner_Call{Call: _e.mock.On("ShardOwner", class, shard)}
}

func (_c *MockshardingStateGetter_ShardOwner_Call) Run(run func(class string, shard string)) *MockshardingStateGetter_ShardOwner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockshardingStateGetter_ShardOwner_Call) Return(_a0 string, _a1 error) *MockshardingStateGetter_ShardOwner_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockshardingStateGetter_ShardOwner_Call) RunAndReturn(run func(string, string) (string, error)) *MockshardingStateGetter_ShardOwner_Call {
	_c.Call.Return(run)
	return _c
}

// ShardReplicas provides a mock function with given fields: class, shard
func (_m *MockshardingStateGetter) ShardReplicas(class string, shard string) ([]string, error) {
	ret := _m.Called(class, shard)

	if len(ret) == 0 {
		panic("no return value specified for ShardReplicas")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, error)); ok {
		return rf(class, shard)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(class, shard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(class, shard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockshardingStateGetter_ShardReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShardReplicas'
type MockshardingStateGetter_ShardReplicas_Call struct {
	*mock.Call
}

// ShardReplicas is a helper method to define mock.On call
//   - class string
//   - shard string
func (_e *MockshardingStateGetter_Expecter) ShardReplicas(class interface{}, shard interface{}) *MockshardingStateGetter_ShardReplicas_Call {
	return &MockshardingStateGetter_ShardReplicas_Call{Call: _e.mock.On("ShardReplicas", class, shard)}
}

func (_c *MockshardingStateGetter_ShardReplicas_Call) Run(run func(class string, shard string)) *MockshardingStateGetter_ShardReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockshardingStateGetter_ShardReplicas_Call) Return(_a0 []string, _a1 error) *MockshardingStateGetter_ShardReplicas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockshardingStateGetter_ShardReplicas_Call) RunAndReturn(run func(string, string) ([]string, error)) *MockshardingStateGetter_ShardReplicas_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockshardingStateGetter creates a new instance of MockshardingStateGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockshardingStateGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockshardingStateGetter {
	mock := &MockshardingStateGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

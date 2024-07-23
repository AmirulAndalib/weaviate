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

import (
	mock "github.com/stretchr/testify/mock"
	schema "github.com/weaviate/weaviate/entities/schema"
)

// MockRemoteIncomingRepo is an autogenerated mock type for the RemoteIncomingRepo type
type MockRemoteIncomingRepo struct {
	mock.Mock
}

type MockRemoteIncomingRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRemoteIncomingRepo) EXPECT() *MockRemoteIncomingRepo_Expecter {
	return &MockRemoteIncomingRepo_Expecter{mock: &_m.Mock}
}

// GetIndexForIncoming provides a mock function with given fields: className
func (_m *MockRemoteIncomingRepo) GetIndexForIncoming(className schema.ClassName) RemoteIndexIncomingRepo {
	ret := _m.Called(className)

	if len(ret) == 0 {
		panic("no return value specified for GetIndexForIncoming")
	}

	var r0 RemoteIndexIncomingRepo
	if rf, ok := ret.Get(0).(func(schema.ClassName) RemoteIndexIncomingRepo); ok {
		r0 = rf(className)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RemoteIndexIncomingRepo)
		}
	}

	return r0
}

// MockRemoteIncomingRepo_GetIndexForIncoming_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndexForIncoming'
type MockRemoteIncomingRepo_GetIndexForIncoming_Call struct {
	*mock.Call
}

// GetIndexForIncoming is a helper method to define mock.On call
//   - className schema.ClassName
func (_e *MockRemoteIncomingRepo_Expecter) GetIndexForIncoming(className interface{}) *MockRemoteIncomingRepo_GetIndexForIncoming_Call {
	return &MockRemoteIncomingRepo_GetIndexForIncoming_Call{Call: _e.mock.On("GetIndexForIncoming", className)}
}

func (_c *MockRemoteIncomingRepo_GetIndexForIncoming_Call) Run(run func(className schema.ClassName)) *MockRemoteIncomingRepo_GetIndexForIncoming_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(schema.ClassName))
	})
	return _c
}

func (_c *MockRemoteIncomingRepo_GetIndexForIncoming_Call) Return(_a0 RemoteIndexIncomingRepo) *MockRemoteIncomingRepo_GetIndexForIncoming_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRemoteIncomingRepo_GetIndexForIncoming_Call) RunAndReturn(run func(schema.ClassName) RemoteIndexIncomingRepo) *MockRemoteIncomingRepo_GetIndexForIncoming_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRemoteIncomingRepo creates a new instance of MockRemoteIncomingRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRemoteIncomingRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRemoteIncomingRepo {
	mock := &MockRemoteIncomingRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

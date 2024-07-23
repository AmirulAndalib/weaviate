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

package locks

import mock "github.com/stretchr/testify/mock"

// MockConnectorSchemaLock is an autogenerated mock type for the ConnectorSchemaLock type
type MockConnectorSchemaLock struct {
	mock.Mock
}

type MockConnectorSchemaLock_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectorSchemaLock) EXPECT() *MockConnectorSchemaLock_Expecter {
	return &MockConnectorSchemaLock_Expecter{mock: &_m.Mock}
}

// LockConnector provides a mock function with given fields:
func (_m *MockConnectorSchemaLock) LockConnector() (func() error, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LockConnector")
	}

	var r0 func() error
	var r1 error
	if rf, ok := ret.Get(0).(func() (func() error, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectorSchemaLock_LockConnector_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LockConnector'
type MockConnectorSchemaLock_LockConnector_Call struct {
	*mock.Call
}

// LockConnector is a helper method to define mock.On call
func (_e *MockConnectorSchemaLock_Expecter) LockConnector() *MockConnectorSchemaLock_LockConnector_Call {
	return &MockConnectorSchemaLock_LockConnector_Call{Call: _e.mock.On("LockConnector")}
}

func (_c *MockConnectorSchemaLock_LockConnector_Call) Run(run func()) *MockConnectorSchemaLock_LockConnector_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConnectorSchemaLock_LockConnector_Call) Return(_a0 func() error, _a1 error) *MockConnectorSchemaLock_LockConnector_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectorSchemaLock_LockConnector_Call) RunAndReturn(run func() (func() error, error)) *MockConnectorSchemaLock_LockConnector_Call {
	_c.Call.Return(run)
	return _c
}

// LockSchema provides a mock function with given fields:
func (_m *MockConnectorSchemaLock) LockSchema() (func() error, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LockSchema")
	}

	var r0 func() error
	var r1 error
	if rf, ok := ret.Get(0).(func() (func() error, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectorSchemaLock_LockSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LockSchema'
type MockConnectorSchemaLock_LockSchema_Call struct {
	*mock.Call
}

// LockSchema is a helper method to define mock.On call
func (_e *MockConnectorSchemaLock_Expecter) LockSchema() *MockConnectorSchemaLock_LockSchema_Call {
	return &MockConnectorSchemaLock_LockSchema_Call{Call: _e.mock.On("LockSchema")}
}

func (_c *MockConnectorSchemaLock_LockSchema_Call) Run(run func()) *MockConnectorSchemaLock_LockSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConnectorSchemaLock_LockSchema_Call) Return(_a0 func() error, _a1 error) *MockConnectorSchemaLock_LockSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectorSchemaLock_LockSchema_Call) RunAndReturn(run func() (func() error, error)) *MockConnectorSchemaLock_LockSchema_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectorSchemaLock creates a new instance of MockConnectorSchemaLock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectorSchemaLock(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectorSchemaLock {
	mock := &MockConnectorSchemaLock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package meta

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"
)

// MockClientOption is an autogenerated mock type for the ClientOption type
type MockClientOption struct {
	mock.Mock
}

type MockClientOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientOption) EXPECT() *MockClientOption_Expecter {
	return &MockClientOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockClientOption) Execute(_a0 *runtime.ClientOperation) {
	_m.Called(_a0)
}

// MockClientOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockClientOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *runtime.ClientOperation
func (_e *MockClientOption_Expecter) Execute(_a0 interface{}) *MockClientOption_Execute_Call {
	return &MockClientOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockClientOption_Execute_Call) Run(run func(_a0 *runtime.ClientOperation)) *MockClientOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*runtime.ClientOperation))
	})
	return _c
}

func (_c *MockClientOption_Execute_Call) Return() *MockClientOption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClientOption_Execute_Call) RunAndReturn(run func(*runtime.ClientOperation)) *MockClientOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientOption creates a new instance of MockClientOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientOption {
	mock := &MockClientOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

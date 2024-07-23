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

package compressionhelpers

import mock "github.com/stretchr/testify/mock"

// MockAction is an autogenerated mock type for the Action type
type MockAction struct {
	mock.Mock
}

type MockAction_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAction) EXPECT() *MockAction_Expecter {
	return &MockAction_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: taskIndex
func (_m *MockAction) Execute(taskIndex uint64) {
	_m.Called(taskIndex)
}

// MockAction_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockAction_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - taskIndex uint64
func (_e *MockAction_Expecter) Execute(taskIndex interface{}) *MockAction_Execute_Call {
	return &MockAction_Execute_Call{Call: _e.mock.On("Execute", taskIndex)}
}

func (_c *MockAction_Execute_Call) Run(run func(taskIndex uint64)) *MockAction_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockAction_Execute_Call) Return() *MockAction_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAction_Execute_Call) RunAndReturn(run func(uint64)) *MockAction_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAction creates a new instance of MockAction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAction(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAction {
	mock := &MockAction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

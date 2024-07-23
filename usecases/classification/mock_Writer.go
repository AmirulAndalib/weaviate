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

package classification

import (
	mock "github.com/stretchr/testify/mock"
	search "github.com/weaviate/weaviate/entities/search"
)

// MockWriter is an autogenerated mock type for the Writer type
type MockWriter struct {
	mock.Mock
}

type MockWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriter) EXPECT() *MockWriter_Expecter {
	return &MockWriter_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields:
func (_m *MockWriter) Start() {
	_m.Called()
}

// MockWriter_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockWriter_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockWriter_Expecter) Start() *MockWriter_Start_Call {
	return &MockWriter_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockWriter_Start_Call) Run(run func()) *MockWriter_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWriter_Start_Call) Return() *MockWriter_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWriter_Start_Call) RunAndReturn(run func()) *MockWriter_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockWriter) Stop() WriterResults {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 WriterResults
	if rf, ok := ret.Get(0).(func() WriterResults); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(WriterResults)
		}
	}

	return r0
}

// MockWriter_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockWriter_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockWriter_Expecter) Stop() *MockWriter_Stop_Call {
	return &MockWriter_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockWriter_Stop_Call) Run(run func()) *MockWriter_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWriter_Stop_Call) Return(_a0 WriterResults) *MockWriter_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWriter_Stop_Call) RunAndReturn(run func() WriterResults) *MockWriter_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Store provides a mock function with given fields: item
func (_m *MockWriter) Store(item search.Result) error {
	ret := _m.Called(item)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(search.Result) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWriter_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type MockWriter_Store_Call struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - item search.Result
func (_e *MockWriter_Expecter) Store(item interface{}) *MockWriter_Store_Call {
	return &MockWriter_Store_Call{Call: _e.mock.On("Store", item)}
}

func (_c *MockWriter_Store_Call) Run(run func(item search.Result)) *MockWriter_Store_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(search.Result))
	})
	return _c
}

func (_c *MockWriter_Store_Call) Return(_a0 error) *MockWriter_Store_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWriter_Store_Call) RunAndReturn(run func(search.Result) error) *MockWriter_Store_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWriter creates a new instance of MockWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWriter {
	mock := &MockWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

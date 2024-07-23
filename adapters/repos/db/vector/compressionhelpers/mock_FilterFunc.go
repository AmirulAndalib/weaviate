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

// MockFilterFunc is an autogenerated mock type for the FilterFunc type
type MockFilterFunc struct {
	mock.Mock
}

type MockFilterFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFilterFunc) EXPECT() *MockFilterFunc_Expecter {
	return &MockFilterFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockFilterFunc) Execute(_a0 []float32) []float32 {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 []float32
	if rf, ok := ret.Get(0).(func([]float32) []float32); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	return r0
}

// MockFilterFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockFilterFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 []float32
func (_e *MockFilterFunc_Expecter) Execute(_a0 interface{}) *MockFilterFunc_Execute_Call {
	return &MockFilterFunc_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockFilterFunc_Execute_Call) Run(run func(_a0 []float32)) *MockFilterFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]float32))
	})
	return _c
}

func (_c *MockFilterFunc_Execute_Call) Return(_a0 []float32) *MockFilterFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilterFunc_Execute_Call) RunAndReturn(run func([]float32) []float32) *MockFilterFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFilterFunc creates a new instance of MockFilterFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFilterFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFilterFunc {
	mock := &MockFilterFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package stopwords

import mock "github.com/stretchr/testify/mock"

// MockStopwordDetector is an autogenerated mock type for the StopwordDetector type
type MockStopwordDetector struct {
	mock.Mock
}

type MockStopwordDetector_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStopwordDetector) EXPECT() *MockStopwordDetector_Expecter {
	return &MockStopwordDetector_Expecter{mock: &_m.Mock}
}

// IsStopword provides a mock function with given fields: _a0
func (_m *MockStopwordDetector) IsStopword(_a0 string) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for IsStopword")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStopwordDetector_IsStopword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsStopword'
type MockStopwordDetector_IsStopword_Call struct {
	*mock.Call
}

// IsStopword is a helper method to define mock.On call
//   - _a0 string
func (_e *MockStopwordDetector_Expecter) IsStopword(_a0 interface{}) *MockStopwordDetector_IsStopword_Call {
	return &MockStopwordDetector_IsStopword_Call{Call: _e.mock.On("IsStopword", _a0)}
}

func (_c *MockStopwordDetector_IsStopword_Call) Run(run func(_a0 string)) *MockStopwordDetector_IsStopword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStopwordDetector_IsStopword_Call) Return(_a0 bool) *MockStopwordDetector_IsStopword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwordDetector_IsStopword_Call) RunAndReturn(run func(string) bool) *MockStopwordDetector_IsStopword_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStopwordDetector creates a new instance of MockStopwordDetector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStopwordDetector(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStopwordDetector {
	mock := &MockStopwordDetector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package vectorizer

import mock "github.com/stretchr/testify/mock"

// MockcalcFn is an autogenerated mock type for the calcFn type
type MockcalcFn struct {
	mock.Mock
}

type MockcalcFn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockcalcFn) EXPECT() *MockcalcFn_Expecter {
	return &MockcalcFn_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: vecs
func (_m *MockcalcFn) Execute(vecs ...[]float32) ([]float32, error) {
	_va := make([]interface{}, len(vecs))
	for _i := range vecs {
		_va[_i] = vecs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 []float32
	var r1 error
	if rf, ok := ret.Get(0).(func(...[]float32) ([]float32, error)); ok {
		return rf(vecs...)
	}
	if rf, ok := ret.Get(0).(func(...[]float32) []float32); ok {
		r0 = rf(vecs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	if rf, ok := ret.Get(1).(func(...[]float32) error); ok {
		r1 = rf(vecs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockcalcFn_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockcalcFn_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - vecs ...[]float32
func (_e *MockcalcFn_Expecter) Execute(vecs ...interface{}) *MockcalcFn_Execute_Call {
	return &MockcalcFn_Execute_Call{Call: _e.mock.On("Execute",
		append([]interface{}{}, vecs...)...)}
}

func (_c *MockcalcFn_Execute_Call) Run(run func(vecs ...[]float32)) *MockcalcFn_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([][]float32, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.([]float32)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockcalcFn_Execute_Call) Return(_a0 []float32, _a1 error) *MockcalcFn_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockcalcFn_Execute_Call) RunAndReturn(run func(...[]float32) ([]float32, error)) *MockcalcFn_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockcalcFn creates a new instance of MockcalcFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockcalcFn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockcalcFn {
	mock := &MockcalcFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

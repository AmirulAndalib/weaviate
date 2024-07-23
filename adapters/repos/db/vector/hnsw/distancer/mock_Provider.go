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

package distancer

import mock "github.com/stretchr/testify/mock"

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

type MockProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProvider) EXPECT() *MockProvider_Expecter {
	return &MockProvider_Expecter{mock: &_m.Mock}
}

// New provides a mock function with given fields: vec
func (_m *MockProvider) New(vec []float32) Distancer {
	ret := _m.Called(vec)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 Distancer
	if rf, ok := ret.Get(0).(func([]float32) Distancer); ok {
		r0 = rf(vec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Distancer)
		}
	}

	return r0
}

// MockProvider_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type MockProvider_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - vec []float32
func (_e *MockProvider_Expecter) New(vec interface{}) *MockProvider_New_Call {
	return &MockProvider_New_Call{Call: _e.mock.On("New", vec)}
}

func (_c *MockProvider_New_Call) Run(run func(vec []float32)) *MockProvider_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]float32))
	})
	return _c
}

func (_c *MockProvider_New_Call) Return(_a0 Distancer) *MockProvider_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProvider_New_Call) RunAndReturn(run func([]float32) Distancer) *MockProvider_New_Call {
	_c.Call.Return(run)
	return _c
}

// SingleDist provides a mock function with given fields: vec1, vec2
func (_m *MockProvider) SingleDist(vec1 []float32, vec2 []float32) (float32, bool, error) {
	ret := _m.Called(vec1, vec2)

	if len(ret) == 0 {
		panic("no return value specified for SingleDist")
	}

	var r0 float32
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func([]float32, []float32) (float32, bool, error)); ok {
		return rf(vec1, vec2)
	}
	if rf, ok := ret.Get(0).(func([]float32, []float32) float32); ok {
		r0 = rf(vec1, vec2)
	} else {
		r0 = ret.Get(0).(float32)
	}

	if rf, ok := ret.Get(1).(func([]float32, []float32) bool); ok {
		r1 = rf(vec1, vec2)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func([]float32, []float32) error); ok {
		r2 = rf(vec1, vec2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockProvider_SingleDist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SingleDist'
type MockProvider_SingleDist_Call struct {
	*mock.Call
}

// SingleDist is a helper method to define mock.On call
//   - vec1 []float32
//   - vec2 []float32
func (_e *MockProvider_Expecter) SingleDist(vec1 interface{}, vec2 interface{}) *MockProvider_SingleDist_Call {
	return &MockProvider_SingleDist_Call{Call: _e.mock.On("SingleDist", vec1, vec2)}
}

func (_c *MockProvider_SingleDist_Call) Run(run func(vec1 []float32, vec2 []float32)) *MockProvider_SingleDist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]float32), args[1].([]float32))
	})
	return _c
}

func (_c *MockProvider_SingleDist_Call) Return(_a0 float32, _a1 bool, _a2 error) *MockProvider_SingleDist_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockProvider_SingleDist_Call) RunAndReturn(run func([]float32, []float32) (float32, bool, error)) *MockProvider_SingleDist_Call {
	_c.Call.Return(run)
	return _c
}

// Step provides a mock function with given fields: x, y
func (_m *MockProvider) Step(x []float32, y []float32) float32 {
	ret := _m.Called(x, y)

	if len(ret) == 0 {
		panic("no return value specified for Step")
	}

	var r0 float32
	if rf, ok := ret.Get(0).(func([]float32, []float32) float32); ok {
		r0 = rf(x, y)
	} else {
		r0 = ret.Get(0).(float32)
	}

	return r0
}

// MockProvider_Step_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Step'
type MockProvider_Step_Call struct {
	*mock.Call
}

// Step is a helper method to define mock.On call
//   - x []float32
//   - y []float32
func (_e *MockProvider_Expecter) Step(x interface{}, y interface{}) *MockProvider_Step_Call {
	return &MockProvider_Step_Call{Call: _e.mock.On("Step", x, y)}
}

func (_c *MockProvider_Step_Call) Run(run func(x []float32, y []float32)) *MockProvider_Step_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]float32), args[1].([]float32))
	})
	return _c
}

func (_c *MockProvider_Step_Call) Return(_a0 float32) *MockProvider_Step_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProvider_Step_Call) RunAndReturn(run func([]float32, []float32) float32) *MockProvider_Step_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *MockProvider) Type() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockProvider_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockProvider_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *MockProvider_Expecter) Type() *MockProvider_Type_Call {
	return &MockProvider_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *MockProvider_Type_Call) Run(run func()) *MockProvider_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_Type_Call) Return(_a0 string) *MockProvider_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProvider_Type_Call) RunAndReturn(run func() string) *MockProvider_Type_Call {
	_c.Call.Return(run)
	return _c
}

// Wrap provides a mock function with given fields: x
func (_m *MockProvider) Wrap(x float32) float32 {
	ret := _m.Called(x)

	if len(ret) == 0 {
		panic("no return value specified for Wrap")
	}

	var r0 float32
	if rf, ok := ret.Get(0).(func(float32) float32); ok {
		r0 = rf(x)
	} else {
		r0 = ret.Get(0).(float32)
	}

	return r0
}

// MockProvider_Wrap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Wrap'
type MockProvider_Wrap_Call struct {
	*mock.Call
}

// Wrap is a helper method to define mock.On call
//   - x float32
func (_e *MockProvider_Expecter) Wrap(x interface{}) *MockProvider_Wrap_Call {
	return &MockProvider_Wrap_Call{Call: _e.mock.On("Wrap", x)}
}

func (_c *MockProvider_Wrap_Call) Run(run func(x float32)) *MockProvider_Wrap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float32))
	})
	return _c
}

func (_c *MockProvider_Wrap_Call) Return(_a0 float32) *MockProvider_Wrap_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProvider_Wrap_Call) RunAndReturn(run func(float32) float32) *MockProvider_Wrap_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProvider creates a new instance of MockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProvider {
	mock := &MockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

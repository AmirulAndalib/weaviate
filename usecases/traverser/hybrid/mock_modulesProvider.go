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

package hybrid

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockmodulesProvider is an autogenerated mock type for the modulesProvider type
type MockmodulesProvider struct {
	mock.Mock
}

type MockmodulesProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockmodulesProvider) EXPECT() *MockmodulesProvider_Expecter {
	return &MockmodulesProvider_Expecter{mock: &_m.Mock}
}

// VectorFromInput provides a mock function with given fields: ctx, className, input, targetVector
func (_m *MockmodulesProvider) VectorFromInput(ctx context.Context, className string, input string, targetVector string) ([]float32, error) {
	ret := _m.Called(ctx, className, input, targetVector)

	if len(ret) == 0 {
		panic("no return value specified for VectorFromInput")
	}

	var r0 []float32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) ([]float32, error)); ok {
		return rf(ctx, className, input, targetVector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []float32); ok {
		r0 = rf(ctx, className, input, targetVector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, className, input, targetVector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockmodulesProvider_VectorFromInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VectorFromInput'
type MockmodulesProvider_VectorFromInput_Call struct {
	*mock.Call
}

// VectorFromInput is a helper method to define mock.On call
//   - ctx context.Context
//   - className string
//   - input string
//   - targetVector string
func (_e *MockmodulesProvider_Expecter) VectorFromInput(ctx interface{}, className interface{}, input interface{}, targetVector interface{}) *MockmodulesProvider_VectorFromInput_Call {
	return &MockmodulesProvider_VectorFromInput_Call{Call: _e.mock.On("VectorFromInput", ctx, className, input, targetVector)}
}

func (_c *MockmodulesProvider_VectorFromInput_Call) Run(run func(ctx context.Context, className string, input string, targetVector string)) *MockmodulesProvider_VectorFromInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockmodulesProvider_VectorFromInput_Call) Return(_a0 []float32, _a1 error) *MockmodulesProvider_VectorFromInput_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockmodulesProvider_VectorFromInput_Call) RunAndReturn(run func(context.Context, string, string, string) ([]float32, error)) *MockmodulesProvider_VectorFromInput_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockmodulesProvider creates a new instance of MockmodulesProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockmodulesProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockmodulesProvider {
	mock := &MockmodulesProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

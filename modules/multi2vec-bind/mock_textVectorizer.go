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

package modbind

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	moduletools "github.com/weaviate/weaviate/entities/moduletools"
)

// MocktextVectorizer is an autogenerated mock type for the textVectorizer type
type MocktextVectorizer struct {
	mock.Mock
}

type MocktextVectorizer_Expecter struct {
	mock *mock.Mock
}

func (_m *MocktextVectorizer) EXPECT() *MocktextVectorizer_Expecter {
	return &MocktextVectorizer_Expecter{mock: &_m.Mock}
}

// Texts provides a mock function with given fields: ctx, input, cfg
func (_m *MocktextVectorizer) Texts(ctx context.Context, input []string, cfg moduletools.ClassConfig) ([]float32, error) {
	ret := _m.Called(ctx, input, cfg)

	if len(ret) == 0 {
		panic("no return value specified for Texts")
	}

	var r0 []float32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, moduletools.ClassConfig) ([]float32, error)); ok {
		return rf(ctx, input, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, moduletools.ClassConfig) []float32); ok {
		r0 = rf(ctx, input, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, moduletools.ClassConfig) error); ok {
		r1 = rf(ctx, input, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MocktextVectorizer_Texts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Texts'
type MocktextVectorizer_Texts_Call struct {
	*mock.Call
}

// Texts is a helper method to define mock.On call
//   - ctx context.Context
//   - input []string
//   - cfg moduletools.ClassConfig
func (_e *MocktextVectorizer_Expecter) Texts(ctx interface{}, input interface{}, cfg interface{}) *MocktextVectorizer_Texts_Call {
	return &MocktextVectorizer_Texts_Call{Call: _e.mock.On("Texts", ctx, input, cfg)}
}

func (_c *MocktextVectorizer_Texts_Call) Run(run func(ctx context.Context, input []string, cfg moduletools.ClassConfig)) *MocktextVectorizer_Texts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MocktextVectorizer_Texts_Call) Return(_a0 []float32, _a1 error) *MocktextVectorizer_Texts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MocktextVectorizer_Texts_Call) RunAndReturn(run func(context.Context, []string, moduletools.ClassConfig) ([]float32, error)) *MocktextVectorizer_Texts_Call {
	_c.Call.Return(run)
	return _c
}

// NewMocktextVectorizer creates a new instance of MocktextVectorizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMocktextVectorizer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MocktextVectorizer {
	mock := &MocktextVectorizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

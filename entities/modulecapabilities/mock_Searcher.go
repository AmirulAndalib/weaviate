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

package modulecapabilities

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	moduletools "github.com/weaviate/weaviate/entities/moduletools"

	strfmt "github.com/go-openapi/strfmt"
)

// MockSearcher is an autogenerated mock type for the Searcher type
type MockSearcher struct {
	mock.Mock
}

type MockSearcher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSearcher) EXPECT() *MockSearcher_Expecter {
	return &MockSearcher_Expecter{mock: &_m.Mock}
}

// VectorSearches provides a mock function with given fields:
func (_m *MockSearcher) VectorSearches() map[string]func(context.Context, interface{}, string, func(context.Context, string, strfmt.UUID, string, string) ([]float32, string, error), moduletools.ClassConfig) ([]float32, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for VectorSearches")
	}

	var r0 map[string]func(context.Context, interface{}, string, func(context.Context, string, strfmt.UUID, string, string) ([]float32, string, error), moduletools.ClassConfig) ([]float32, error)
	if rf, ok := ret.Get(0).(func() map[string]func(context.Context, interface{}, string, func(context.Context, string, strfmt.UUID, string, string) ([]float32, string, error), moduletools.ClassConfig) ([]float32, error)); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]func(context.Context, interface{}, string, func(context.Context, string, strfmt.UUID, string, string) ([]float32, string, error), moduletools.ClassConfig) ([]float32, error))
		}
	}

	return r0
}

// MockSearcher_VectorSearches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VectorSearches'
type MockSearcher_VectorSearches_Call struct {
	*mock.Call
}

// VectorSearches is a helper method to define mock.On call
func (_e *MockSearcher_Expecter) VectorSearches() *MockSearcher_VectorSearches_Call {
	return &MockSearcher_VectorSearches_Call{Call: _e.mock.On("VectorSearches")}
}

func (_c *MockSearcher_VectorSearches_Call) Run(run func()) *MockSearcher_VectorSearches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSearcher_VectorSearches_Call) Return(_a0 map[string]func(context.Context, interface{}, string, func(context.Context, string, strfmt.UUID, string, string) ([]float32, string, error), moduletools.ClassConfig) ([]float32, error)) *MockSearcher_VectorSearches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSearcher_VectorSearches_Call) RunAndReturn(run func() map[string]func(context.Context, interface{}, string, func(context.Context, string, strfmt.UUID, string, string) ([]float32, string, error), moduletools.ClassConfig) ([]float32, error)) *MockSearcher_VectorSearches_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSearcher creates a new instance of MockSearcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSearcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSearcher {
	mock := &MockSearcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	ent "github.com/weaviate/weaviate/modules/text2vec-jinaai/ent"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// Vectorize provides a mock function with given fields: ctx, input, config
func (_m *MockClient) Vectorize(ctx context.Context, input string, config ent.VectorizationConfig) (*ent.VectorizationResult, error) {
	ret := _m.Called(ctx, input, config)

	if len(ret) == 0 {
		panic("no return value specified for Vectorize")
	}

	var r0 *ent.VectorizationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ent.VectorizationConfig) (*ent.VectorizationResult, error)); ok {
		return rf(ctx, input, config)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ent.VectorizationConfig) *ent.VectorizationResult); ok {
		r0 = rf(ctx, input, config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.VectorizationResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ent.VectorizationConfig) error); ok {
		r1 = rf(ctx, input, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_Vectorize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Vectorize'
type MockClient_Vectorize_Call struct {
	*mock.Call
}

// Vectorize is a helper method to define mock.On call
//   - ctx context.Context
//   - input string
//   - config ent.VectorizationConfig
func (_e *MockClient_Expecter) Vectorize(ctx interface{}, input interface{}, config interface{}) *MockClient_Vectorize_Call {
	return &MockClient_Vectorize_Call{Call: _e.mock.On("Vectorize", ctx, input, config)}
}

func (_c *MockClient_Vectorize_Call) Run(run func(ctx context.Context, input string, config ent.VectorizationConfig)) *MockClient_Vectorize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(ent.VectorizationConfig))
	})
	return _c
}

func (_c *MockClient_Vectorize_Call) Return(_a0 *ent.VectorizationResult, _a1 error) *MockClient_Vectorize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_Vectorize_Call) RunAndReturn(run func(context.Context, string, ent.VectorizationConfig) (*ent.VectorizationResult, error)) *MockClient_Vectorize_Call {
	_c.Call.Return(run)
	return _c
}

// VectorizeQuery provides a mock function with given fields: ctx, input, config
func (_m *MockClient) VectorizeQuery(ctx context.Context, input []string, config ent.VectorizationConfig) (*ent.VectorizationResult, error) {
	ret := _m.Called(ctx, input, config)

	if len(ret) == 0 {
		panic("no return value specified for VectorizeQuery")
	}

	var r0 *ent.VectorizationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, ent.VectorizationConfig) (*ent.VectorizationResult, error)); ok {
		return rf(ctx, input, config)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, ent.VectorizationConfig) *ent.VectorizationResult); ok {
		r0 = rf(ctx, input, config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.VectorizationResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, ent.VectorizationConfig) error); ok {
		r1 = rf(ctx, input, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_VectorizeQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VectorizeQuery'
type MockClient_VectorizeQuery_Call struct {
	*mock.Call
}

// VectorizeQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - input []string
//   - config ent.VectorizationConfig
func (_e *MockClient_Expecter) VectorizeQuery(ctx interface{}, input interface{}, config interface{}) *MockClient_VectorizeQuery_Call {
	return &MockClient_VectorizeQuery_Call{Call: _e.mock.On("VectorizeQuery", ctx, input, config)}
}

func (_c *MockClient_VectorizeQuery_Call) Run(run func(ctx context.Context, input []string, config ent.VectorizationConfig)) *MockClient_VectorizeQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(ent.VectorizationConfig))
	})
	return _c
}

func (_c *MockClient_VectorizeQuery_Call) Return(_a0 *ent.VectorizationResult, _a1 error) *MockClient_VectorizeQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_VectorizeQuery_Call) RunAndReturn(run func(context.Context, []string, ent.VectorizationConfig) (*ent.VectorizationResult, error)) *MockClient_VectorizeQuery_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

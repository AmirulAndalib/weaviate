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
	ent "github.com/weaviate/weaviate/modules/text2vec-gpt4all/ent"
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

// Vectorize provides a mock function with given fields: ctx, text
func (_m *MockClient) Vectorize(ctx context.Context, text string) (*ent.VectorizationResult, error) {
	ret := _m.Called(ctx, text)

	if len(ret) == 0 {
		panic("no return value specified for Vectorize")
	}

	var r0 *ent.VectorizationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.VectorizationResult, error)); ok {
		return rf(ctx, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.VectorizationResult); ok {
		r0 = rf(ctx, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.VectorizationResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, text)
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
//   - text string
func (_e *MockClient_Expecter) Vectorize(ctx interface{}, text interface{}) *MockClient_Vectorize_Call {
	return &MockClient_Vectorize_Call{Call: _e.mock.On("Vectorize", ctx, text)}
}

func (_c *MockClient_Vectorize_Call) Run(run func(ctx context.Context, text string)) *MockClient_Vectorize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockClient_Vectorize_Call) Return(_a0 *ent.VectorizationResult, _a1 error) *MockClient_Vectorize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_Vectorize_Call) RunAndReturn(run func(context.Context, string) (*ent.VectorizationResult, error)) *MockClient_Vectorize_Call {
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

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

package answer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	ent "github.com/weaviate/weaviate/modules/qna-openai/ent"

	moduletools "github.com/weaviate/weaviate/entities/moduletools"
)

// MockqnaClient is an autogenerated mock type for the qnaClient type
type MockqnaClient struct {
	mock.Mock
}

type MockqnaClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockqnaClient) EXPECT() *MockqnaClient_Expecter {
	return &MockqnaClient_Expecter{mock: &_m.Mock}
}

// Answer provides a mock function with given fields: ctx, text, question, cfg
func (_m *MockqnaClient) Answer(ctx context.Context, text string, question string, cfg moduletools.ClassConfig) (*ent.AnswerResult, error) {
	ret := _m.Called(ctx, text, question, cfg)

	if len(ret) == 0 {
		panic("no return value specified for Answer")
	}

	var r0 *ent.AnswerResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, moduletools.ClassConfig) (*ent.AnswerResult, error)); ok {
		return rf(ctx, text, question, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, moduletools.ClassConfig) *ent.AnswerResult); ok {
		r0 = rf(ctx, text, question, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.AnswerResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, moduletools.ClassConfig) error); ok {
		r1 = rf(ctx, text, question, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockqnaClient_Answer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Answer'
type MockqnaClient_Answer_Call struct {
	*mock.Call
}

// Answer is a helper method to define mock.On call
//   - ctx context.Context
//   - text string
//   - question string
//   - cfg moduletools.ClassConfig
func (_e *MockqnaClient_Expecter) Answer(ctx interface{}, text interface{}, question interface{}, cfg interface{}) *MockqnaClient_Answer_Call {
	return &MockqnaClient_Answer_Call{Call: _e.mock.On("Answer", ctx, text, question, cfg)}
}

func (_c *MockqnaClient_Answer_Call) Run(run func(ctx context.Context, text string, question string, cfg moduletools.ClassConfig)) *MockqnaClient_Answer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MockqnaClient_Answer_Call) Return(_a0 *ent.AnswerResult, _a1 error) *MockqnaClient_Answer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockqnaClient_Answer_Call) RunAndReturn(run func(context.Context, string, string, moduletools.ClassConfig) (*ent.AnswerResult, error)) *MockqnaClient_Answer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockqnaClient creates a new instance of MockqnaClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockqnaClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockqnaClient {
	mock := &MockqnaClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

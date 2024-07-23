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

package additional

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	ent "github.com/weaviate/weaviate/usecases/modulecomponents/ent"

	moduletools "github.com/weaviate/weaviate/entities/moduletools"
)

// MockreRankerClient is an autogenerated mock type for the reRankerClient type
type MockreRankerClient struct {
	mock.Mock
}

type MockreRankerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockreRankerClient) EXPECT() *MockreRankerClient_Expecter {
	return &MockreRankerClient_Expecter{mock: &_m.Mock}
}

// Rank provides a mock function with given fields: ctx, query, documents, cfg
func (_m *MockreRankerClient) Rank(ctx context.Context, query string, documents []string, cfg moduletools.ClassConfig) (*ent.RankResult, error) {
	ret := _m.Called(ctx, query, documents, cfg)

	if len(ret) == 0 {
		panic("no return value specified for Rank")
	}

	var r0 *ent.RankResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, moduletools.ClassConfig) (*ent.RankResult, error)); ok {
		return rf(ctx, query, documents, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, moduletools.ClassConfig) *ent.RankResult); ok {
		r0 = rf(ctx, query, documents, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.RankResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string, moduletools.ClassConfig) error); ok {
		r1 = rf(ctx, query, documents, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockreRankerClient_Rank_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rank'
type MockreRankerClient_Rank_Call struct {
	*mock.Call
}

// Rank is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - documents []string
//   - cfg moduletools.ClassConfig
func (_e *MockreRankerClient_Expecter) Rank(ctx interface{}, query interface{}, documents interface{}, cfg interface{}) *MockreRankerClient_Rank_Call {
	return &MockreRankerClient_Rank_Call{Call: _e.mock.On("Rank", ctx, query, documents, cfg)}
}

func (_c *MockreRankerClient_Rank_Call) Run(run func(ctx context.Context, query string, documents []string, cfg moduletools.ClassConfig)) *MockreRankerClient_Rank_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string), args[3].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MockreRankerClient_Rank_Call) Return(_a0 *ent.RankResult, _a1 error) *MockreRankerClient_Rank_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockreRankerClient_Rank_Call) RunAndReturn(run func(context.Context, string, []string, moduletools.ClassConfig) (*ent.RankResult, error)) *MockreRankerClient_Rank_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockreRankerClient creates a new instance of MockreRankerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockreRankerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockreRankerClient {
	mock := &MockreRankerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

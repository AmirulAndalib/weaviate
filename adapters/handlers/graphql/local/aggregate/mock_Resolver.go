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

package aggregate

import (
	context "context"

	aggregation "github.com/weaviate/weaviate/entities/aggregation"

	mock "github.com/stretchr/testify/mock"

	models "github.com/weaviate/weaviate/entities/models"
)

// MockResolver is an autogenerated mock type for the Resolver type
type MockResolver struct {
	mock.Mock
}

type MockResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResolver) EXPECT() *MockResolver_Expecter {
	return &MockResolver_Expecter{mock: &_m.Mock}
}

// Aggregate provides a mock function with given fields: ctx, principal, info
func (_m *MockResolver) Aggregate(ctx context.Context, principal *models.Principal, info *aggregation.Params) (interface{}, error) {
	ret := _m.Called(ctx, principal, info)

	if len(ret) == 0 {
		panic("no return value specified for Aggregate")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Principal, *aggregation.Params) (interface{}, error)); ok {
		return rf(ctx, principal, info)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Principal, *aggregation.Params) interface{}); ok {
		r0 = rf(ctx, principal, info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Principal, *aggregation.Params) error); ok {
		r1 = rf(ctx, principal, info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResolver_Aggregate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Aggregate'
type MockResolver_Aggregate_Call struct {
	*mock.Call
}

// Aggregate is a helper method to define mock.On call
//   - ctx context.Context
//   - principal *models.Principal
//   - info *aggregation.Params
func (_e *MockResolver_Expecter) Aggregate(ctx interface{}, principal interface{}, info interface{}) *MockResolver_Aggregate_Call {
	return &MockResolver_Aggregate_Call{Call: _e.mock.On("Aggregate", ctx, principal, info)}
}

func (_c *MockResolver_Aggregate_Call) Run(run func(ctx context.Context, principal *models.Principal, info *aggregation.Params)) *MockResolver_Aggregate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Principal), args[2].(*aggregation.Params))
	})
	return _c
}

func (_c *MockResolver_Aggregate_Call) Return(_a0 interface{}, _a1 error) *MockResolver_Aggregate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResolver_Aggregate_Call) RunAndReturn(run func(context.Context, *models.Principal, *aggregation.Params) (interface{}, error)) *MockResolver_Aggregate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockResolver creates a new instance of MockResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResolver(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResolver {
	mock := &MockResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

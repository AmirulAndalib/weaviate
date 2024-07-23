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

package traverser

import (
	context "context"

	aggregation "github.com/weaviate/weaviate/entities/aggregation"

	dto "github.com/weaviate/weaviate/entities/dto"

	mock "github.com/stretchr/testify/mock"
)

// MockTraverserRepo is an autogenerated mock type for the TraverserRepo type
type MockTraverserRepo struct {
	mock.Mock
}

type MockTraverserRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTraverserRepo) EXPECT() *MockTraverserRepo_Expecter {
	return &MockTraverserRepo_Expecter{mock: &_m.Mock}
}

// Aggregate provides a mock function with given fields: _a0, _a1
func (_m *MockTraverserRepo) Aggregate(_a0 context.Context, _a1 *aggregation.Params) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Aggregate")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregation.Params) (interface{}, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *aggregation.Params) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *aggregation.Params) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTraverserRepo_Aggregate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Aggregate'
type MockTraverserRepo_Aggregate_Call struct {
	*mock.Call
}

// Aggregate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *aggregation.Params
func (_e *MockTraverserRepo_Expecter) Aggregate(_a0 interface{}, _a1 interface{}) *MockTraverserRepo_Aggregate_Call {
	return &MockTraverserRepo_Aggregate_Call{Call: _e.mock.On("Aggregate", _a0, _a1)}
}

func (_c *MockTraverserRepo_Aggregate_Call) Run(run func(_a0 context.Context, _a1 *aggregation.Params)) *MockTraverserRepo_Aggregate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*aggregation.Params))
	})
	return _c
}

func (_c *MockTraverserRepo_Aggregate_Call) Return(_a0 interface{}, _a1 error) *MockTraverserRepo_Aggregate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTraverserRepo_Aggregate_Call) RunAndReturn(run func(context.Context, *aggregation.Params) (interface{}, error)) *MockTraverserRepo_Aggregate_Call {
	_c.Call.Return(run)
	return _c
}

// GetClass provides a mock function with given fields: _a0, _a1
func (_m *MockTraverserRepo) GetClass(_a0 context.Context, _a1 *dto.GetParams) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetClass")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetParams) (interface{}, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetParams) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTraverserRepo_GetClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClass'
type MockTraverserRepo_GetClass_Call struct {
	*mock.Call
}

// GetClass is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *dto.GetParams
func (_e *MockTraverserRepo_Expecter) GetClass(_a0 interface{}, _a1 interface{}) *MockTraverserRepo_GetClass_Call {
	return &MockTraverserRepo_GetClass_Call{Call: _e.mock.On("GetClass", _a0, _a1)}
}

func (_c *MockTraverserRepo_GetClass_Call) Run(run func(_a0 context.Context, _a1 *dto.GetParams)) *MockTraverserRepo_GetClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*dto.GetParams))
	})
	return _c
}

func (_c *MockTraverserRepo_GetClass_Call) Return(_a0 interface{}, _a1 error) *MockTraverserRepo_GetClass_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTraverserRepo_GetClass_Call) RunAndReturn(run func(context.Context, *dto.GetParams) (interface{}, error)) *MockTraverserRepo_GetClass_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTraverserRepo creates a new instance of MockTraverserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTraverserRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTraverserRepo {
	mock := &MockTraverserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

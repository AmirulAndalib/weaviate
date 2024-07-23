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

package sorter

import (
	context "context"

	helpers "github.com/weaviate/weaviate/adapters/repos/db/helpers"
	filters "github.com/weaviate/weaviate/entities/filters"

	mock "github.com/stretchr/testify/mock"
)

// MockLSMSorter is an autogenerated mock type for the LSMSorter type
type MockLSMSorter struct {
	mock.Mock
}

type MockLSMSorter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLSMSorter) EXPECT() *MockLSMSorter_Expecter {
	return &MockLSMSorter_Expecter{mock: &_m.Mock}
}

// Sort provides a mock function with given fields: ctx, limit, sort
func (_m *MockLSMSorter) Sort(ctx context.Context, limit int, sort []filters.Sort) ([]uint64, error) {
	ret := _m.Called(ctx, limit, sort)

	if len(ret) == 0 {
		panic("no return value specified for Sort")
	}

	var r0 []uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []filters.Sort) ([]uint64, error)); ok {
		return rf(ctx, limit, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, []filters.Sort) []uint64); ok {
		r0 = rf(ctx, limit, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, []filters.Sort) error); ok {
		r1 = rf(ctx, limit, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLSMSorter_Sort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sort'
type MockLSMSorter_Sort_Call struct {
	*mock.Call
}

// Sort is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int
//   - sort []filters.Sort
func (_e *MockLSMSorter_Expecter) Sort(ctx interface{}, limit interface{}, sort interface{}) *MockLSMSorter_Sort_Call {
	return &MockLSMSorter_Sort_Call{Call: _e.mock.On("Sort", ctx, limit, sort)}
}

func (_c *MockLSMSorter_Sort_Call) Run(run func(ctx context.Context, limit int, sort []filters.Sort)) *MockLSMSorter_Sort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].([]filters.Sort))
	})
	return _c
}

func (_c *MockLSMSorter_Sort_Call) Return(_a0 []uint64, _a1 error) *MockLSMSorter_Sort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLSMSorter_Sort_Call) RunAndReturn(run func(context.Context, int, []filters.Sort) ([]uint64, error)) *MockLSMSorter_Sort_Call {
	_c.Call.Return(run)
	return _c
}

// SortDocIDs provides a mock function with given fields: ctx, limit, sort, ids
func (_m *MockLSMSorter) SortDocIDs(ctx context.Context, limit int, sort []filters.Sort, ids helpers.AllowList) ([]uint64, error) {
	ret := _m.Called(ctx, limit, sort, ids)

	if len(ret) == 0 {
		panic("no return value specified for SortDocIDs")
	}

	var r0 []uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []filters.Sort, helpers.AllowList) ([]uint64, error)); ok {
		return rf(ctx, limit, sort, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, []filters.Sort, helpers.AllowList) []uint64); ok {
		r0 = rf(ctx, limit, sort, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, []filters.Sort, helpers.AllowList) error); ok {
		r1 = rf(ctx, limit, sort, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLSMSorter_SortDocIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SortDocIDs'
type MockLSMSorter_SortDocIDs_Call struct {
	*mock.Call
}

// SortDocIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int
//   - sort []filters.Sort
//   - ids helpers.AllowList
func (_e *MockLSMSorter_Expecter) SortDocIDs(ctx interface{}, limit interface{}, sort interface{}, ids interface{}) *MockLSMSorter_SortDocIDs_Call {
	return &MockLSMSorter_SortDocIDs_Call{Call: _e.mock.On("SortDocIDs", ctx, limit, sort, ids)}
}

func (_c *MockLSMSorter_SortDocIDs_Call) Run(run func(ctx context.Context, limit int, sort []filters.Sort, ids helpers.AllowList)) *MockLSMSorter_SortDocIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].([]filters.Sort), args[3].(helpers.AllowList))
	})
	return _c
}

func (_c *MockLSMSorter_SortDocIDs_Call) Return(_a0 []uint64, _a1 error) *MockLSMSorter_SortDocIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLSMSorter_SortDocIDs_Call) RunAndReturn(run func(context.Context, int, []filters.Sort, helpers.AllowList) ([]uint64, error)) *MockLSMSorter_SortDocIDs_Call {
	_c.Call.Return(run)
	return _c
}

// SortDocIDsAndDists provides a mock function with given fields: ctx, limit, sort, ids, dists
func (_m *MockLSMSorter) SortDocIDsAndDists(ctx context.Context, limit int, sort []filters.Sort, ids []uint64, dists []float32) ([]uint64, []float32, error) {
	ret := _m.Called(ctx, limit, sort, ids, dists)

	if len(ret) == 0 {
		panic("no return value specified for SortDocIDsAndDists")
	}

	var r0 []uint64
	var r1 []float32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []filters.Sort, []uint64, []float32) ([]uint64, []float32, error)); ok {
		return rf(ctx, limit, sort, ids, dists)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, []filters.Sort, []uint64, []float32) []uint64); ok {
		r0 = rf(ctx, limit, sort, ids, dists)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, []filters.Sort, []uint64, []float32) []float32); ok {
		r1 = rf(ctx, limit, sort, ids, dists)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]float32)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, []filters.Sort, []uint64, []float32) error); ok {
		r2 = rf(ctx, limit, sort, ids, dists)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockLSMSorter_SortDocIDsAndDists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SortDocIDsAndDists'
type MockLSMSorter_SortDocIDsAndDists_Call struct {
	*mock.Call
}

// SortDocIDsAndDists is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int
//   - sort []filters.Sort
//   - ids []uint64
//   - dists []float32
func (_e *MockLSMSorter_Expecter) SortDocIDsAndDists(ctx interface{}, limit interface{}, sort interface{}, ids interface{}, dists interface{}) *MockLSMSorter_SortDocIDsAndDists_Call {
	return &MockLSMSorter_SortDocIDsAndDists_Call{Call: _e.mock.On("SortDocIDsAndDists", ctx, limit, sort, ids, dists)}
}

func (_c *MockLSMSorter_SortDocIDsAndDists_Call) Run(run func(ctx context.Context, limit int, sort []filters.Sort, ids []uint64, dists []float32)) *MockLSMSorter_SortDocIDsAndDists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].([]filters.Sort), args[3].([]uint64), args[4].([]float32))
	})
	return _c
}

func (_c *MockLSMSorter_SortDocIDsAndDists_Call) Return(_a0 []uint64, _a1 []float32, _a2 error) *MockLSMSorter_SortDocIDsAndDists_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockLSMSorter_SortDocIDsAndDists_Call) RunAndReturn(run func(context.Context, int, []filters.Sort, []uint64, []float32) ([]uint64, []float32, error)) *MockLSMSorter_SortDocIDsAndDists_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLSMSorter creates a new instance of MockLSMSorter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLSMSorter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLSMSorter {
	mock := &MockLSMSorter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

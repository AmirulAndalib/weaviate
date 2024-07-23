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

package cache

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache[T interface{}] struct {
	mock.Mock
}

type MockCache_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockCache[T]) EXPECT() *MockCache_Expecter[T] {
	return &MockCache_Expecter[T]{mock: &_m.Mock}
}

// All provides a mock function with given fields:
func (_m *MockCache[T]) All() [][]T {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 [][]T
	if rf, ok := ret.Get(0).(func() [][]T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]T)
		}
	}

	return r0
}

// MockCache_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type MockCache_All_Call[T interface{}] struct {
	*mock.Call
}

// All is a helper method to define mock.On call
func (_e *MockCache_Expecter[T]) All() *MockCache_All_Call[T] {
	return &MockCache_All_Call[T]{Call: _e.mock.On("All")}
}

func (_c *MockCache_All_Call[T]) Run(run func()) *MockCache_All_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCache_All_Call[T]) Return(_a0 [][]T) *MockCache_All_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_All_Call[T]) RunAndReturn(run func() [][]T) *MockCache_All_Call[T] {
	_c.Call.Return(run)
	return _c
}

// CopyMaxSize provides a mock function with given fields:
func (_m *MockCache[T]) CopyMaxSize() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CopyMaxSize")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockCache_CopyMaxSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyMaxSize'
type MockCache_CopyMaxSize_Call[T interface{}] struct {
	*mock.Call
}

// CopyMaxSize is a helper method to define mock.On call
func (_e *MockCache_Expecter[T]) CopyMaxSize() *MockCache_CopyMaxSize_Call[T] {
	return &MockCache_CopyMaxSize_Call[T]{Call: _e.mock.On("CopyMaxSize")}
}

func (_c *MockCache_CopyMaxSize_Call[T]) Run(run func()) *MockCache_CopyMaxSize_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCache_CopyMaxSize_Call[T]) Return(_a0 int64) *MockCache_CopyMaxSize_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_CopyMaxSize_Call[T]) RunAndReturn(run func() int64) *MockCache_CopyMaxSize_Call[T] {
	_c.Call.Return(run)
	return _c
}

// CountVectors provides a mock function with given fields:
func (_m *MockCache[T]) CountVectors() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CountVectors")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockCache_CountVectors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountVectors'
type MockCache_CountVectors_Call[T interface{}] struct {
	*mock.Call
}

// CountVectors is a helper method to define mock.On call
func (_e *MockCache_Expecter[T]) CountVectors() *MockCache_CountVectors_Call[T] {
	return &MockCache_CountVectors_Call[T]{Call: _e.mock.On("CountVectors")}
}

func (_c *MockCache_CountVectors_Call[T]) Run(run func()) *MockCache_CountVectors_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCache_CountVectors_Call[T]) Return(_a0 int64) *MockCache_CountVectors_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_CountVectors_Call[T]) RunAndReturn(run func() int64) *MockCache_CountVectors_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockCache[T]) Delete(ctx context.Context, id uint64) {
	_m.Called(ctx, id)
}

// MockCache_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCache_Delete_Call[T interface{}] struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockCache_Expecter[T]) Delete(ctx interface{}, id interface{}) *MockCache_Delete_Call[T] {
	return &MockCache_Delete_Call[T]{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockCache_Delete_Call[T]) Run(run func(ctx context.Context, id uint64)) *MockCache_Delete_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockCache_Delete_Call[T]) Return() *MockCache_Delete_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_Delete_Call[T]) RunAndReturn(run func(context.Context, uint64)) *MockCache_Delete_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Drop provides a mock function with given fields:
func (_m *MockCache[T]) Drop() {
	_m.Called()
}

// MockCache_Drop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drop'
type MockCache_Drop_Call[T interface{}] struct {
	*mock.Call
}

// Drop is a helper method to define mock.On call
func (_e *MockCache_Expecter[T]) Drop() *MockCache_Drop_Call[T] {
	return &MockCache_Drop_Call[T]{Call: _e.mock.On("Drop")}
}

func (_c *MockCache_Drop_Call[T]) Run(run func()) *MockCache_Drop_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCache_Drop_Call[T]) Return() *MockCache_Drop_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_Drop_Call[T]) RunAndReturn(run func()) *MockCache_Drop_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockCache[T]) Get(ctx context.Context, id uint64) ([]T, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]T, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []T); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCache_Get_Call[T interface{}] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint64
func (_e *MockCache_Expecter[T]) Get(ctx interface{}, id interface{}) *MockCache_Get_Call[T] {
	return &MockCache_Get_Call[T]{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockCache_Get_Call[T]) Run(run func(ctx context.Context, id uint64)) *MockCache_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockCache_Get_Call[T]) Return(_a0 []T, _a1 error) *MockCache_Get_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_Get_Call[T]) RunAndReturn(run func(context.Context, uint64) ([]T, error)) *MockCache_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Grow provides a mock function with given fields: size
func (_m *MockCache[T]) Grow(size uint64) {
	_m.Called(size)
}

// MockCache_Grow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Grow'
type MockCache_Grow_Call[T interface{}] struct {
	*mock.Call
}

// Grow is a helper method to define mock.On call
//   - size uint64
func (_e *MockCache_Expecter[T]) Grow(size interface{}) *MockCache_Grow_Call[T] {
	return &MockCache_Grow_Call[T]{Call: _e.mock.On("Grow", size)}
}

func (_c *MockCache_Grow_Call[T]) Run(run func(size uint64)) *MockCache_Grow_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockCache_Grow_Call[T]) Return() *MockCache_Grow_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_Grow_Call[T]) RunAndReturn(run func(uint64)) *MockCache_Grow_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Len provides a mock function with given fields:
func (_m *MockCache[T]) Len() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Len")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// MockCache_Len_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Len'
type MockCache_Len_Call[T interface{}] struct {
	*mock.Call
}

// Len is a helper method to define mock.On call
func (_e *MockCache_Expecter[T]) Len() *MockCache_Len_Call[T] {
	return &MockCache_Len_Call[T]{Call: _e.mock.On("Len")}
}

func (_c *MockCache_Len_Call[T]) Run(run func()) *MockCache_Len_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCache_Len_Call[T]) Return(_a0 int32) *MockCache_Len_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_Len_Call[T]) RunAndReturn(run func() int32) *MockCache_Len_Call[T] {
	_c.Call.Return(run)
	return _c
}

// MultiGet provides a mock function with given fields: ctx, ids
func (_m *MockCache[T]) MultiGet(ctx context.Context, ids []uint64) ([][]T, []error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for MultiGet")
	}

	var r0 [][]T
	var r1 []error
	if rf, ok := ret.Get(0).(func(context.Context, []uint64) ([][]T, []error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint64) [][]T); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint64) []error); ok {
		r1 = rf(ctx, ids)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// MockCache_MultiGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiGet'
type MockCache_MultiGet_Call[T interface{}] struct {
	*mock.Call
}

// MultiGet is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uint64
func (_e *MockCache_Expecter[T]) MultiGet(ctx interface{}, ids interface{}) *MockCache_MultiGet_Call[T] {
	return &MockCache_MultiGet_Call[T]{Call: _e.mock.On("MultiGet", ctx, ids)}
}

func (_c *MockCache_MultiGet_Call[T]) Run(run func(ctx context.Context, ids []uint64)) *MockCache_MultiGet_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint64))
	})
	return _c
}

func (_c *MockCache_MultiGet_Call[T]) Return(_a0 [][]T, _a1 []error) *MockCache_MultiGet_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_MultiGet_Call[T]) RunAndReturn(run func(context.Context, []uint64) ([][]T, []error)) *MockCache_MultiGet_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Prefetch provides a mock function with given fields: id
func (_m *MockCache[T]) Prefetch(id uint64) {
	_m.Called(id)
}

// MockCache_Prefetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prefetch'
type MockCache_Prefetch_Call[T interface{}] struct {
	*mock.Call
}

// Prefetch is a helper method to define mock.On call
//   - id uint64
func (_e *MockCache_Expecter[T]) Prefetch(id interface{}) *MockCache_Prefetch_Call[T] {
	return &MockCache_Prefetch_Call[T]{Call: _e.mock.On("Prefetch", id)}
}

func (_c *MockCache_Prefetch_Call[T]) Run(run func(id uint64)) *MockCache_Prefetch_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockCache_Prefetch_Call[T]) Return() *MockCache_Prefetch_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_Prefetch_Call[T]) RunAndReturn(run func(uint64)) *MockCache_Prefetch_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Preload provides a mock function with given fields: id, vec
func (_m *MockCache[T]) Preload(id uint64, vec []T) {
	_m.Called(id, vec)
}

// MockCache_Preload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Preload'
type MockCache_Preload_Call[T interface{}] struct {
	*mock.Call
}

// Preload is a helper method to define mock.On call
//   - id uint64
//   - vec []T
func (_e *MockCache_Expecter[T]) Preload(id interface{}, vec interface{}) *MockCache_Preload_Call[T] {
	return &MockCache_Preload_Call[T]{Call: _e.mock.On("Preload", id, vec)}
}

func (_c *MockCache_Preload_Call[T]) Run(run func(id uint64, vec []T)) *MockCache_Preload_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].([]T))
	})
	return _c
}

func (_c *MockCache_Preload_Call[T]) Return() *MockCache_Preload_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_Preload_Call[T]) RunAndReturn(run func(uint64, []T)) *MockCache_Preload_Call[T] {
	_c.Call.Return(run)
	return _c
}

// UpdateMaxSize provides a mock function with given fields: size
func (_m *MockCache[T]) UpdateMaxSize(size int64) {
	_m.Called(size)
}

// MockCache_UpdateMaxSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMaxSize'
type MockCache_UpdateMaxSize_Call[T interface{}] struct {
	*mock.Call
}

// UpdateMaxSize is a helper method to define mock.On call
//   - size int64
func (_e *MockCache_Expecter[T]) UpdateMaxSize(size interface{}) *MockCache_UpdateMaxSize_Call[T] {
	return &MockCache_UpdateMaxSize_Call[T]{Call: _e.mock.On("UpdateMaxSize", size)}
}

func (_c *MockCache_UpdateMaxSize_Call[T]) Run(run func(size int64)) *MockCache_UpdateMaxSize_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockCache_UpdateMaxSize_Call[T]) Return() *MockCache_UpdateMaxSize_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_UpdateMaxSize_Call[T]) RunAndReturn(run func(int64)) *MockCache_UpdateMaxSize_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCache[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCache[T] {
	mock := &MockCache[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

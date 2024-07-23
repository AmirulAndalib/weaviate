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

package classifications

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/weaviate/weaviate/entities/models"

	strfmt "github.com/go-openapi/strfmt"
)

// MocklocalRepo is an autogenerated mock type for the localRepo type
type MocklocalRepo struct {
	mock.Mock
}

type MocklocalRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MocklocalRepo) EXPECT() *MocklocalRepo_Expecter {
	return &MocklocalRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *MocklocalRepo) Get(ctx context.Context, id strfmt.UUID) (*models.Classification, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Classification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, strfmt.UUID) (*models.Classification, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, strfmt.UUID) *models.Classification); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Classification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, strfmt.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MocklocalRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MocklocalRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id strfmt.UUID
func (_e *MocklocalRepo_Expecter) Get(ctx interface{}, id interface{}) *MocklocalRepo_Get_Call {
	return &MocklocalRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MocklocalRepo_Get_Call) Run(run func(ctx context.Context, id strfmt.UUID)) *MocklocalRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(strfmt.UUID))
	})
	return _c
}

func (_c *MocklocalRepo_Get_Call) Return(_a0 *models.Classification, _a1 error) *MocklocalRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MocklocalRepo_Get_Call) RunAndReturn(run func(context.Context, strfmt.UUID) (*models.Classification, error)) *MocklocalRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, classification
func (_m *MocklocalRepo) Put(ctx context.Context, classification models.Classification) error {
	ret := _m.Called(ctx, classification)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Classification) error); ok {
		r0 = rf(ctx, classification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MocklocalRepo_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MocklocalRepo_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - classification models.Classification
func (_e *MocklocalRepo_Expecter) Put(ctx interface{}, classification interface{}) *MocklocalRepo_Put_Call {
	return &MocklocalRepo_Put_Call{Call: _e.mock.On("Put", ctx, classification)}
}

func (_c *MocklocalRepo_Put_Call) Run(run func(ctx context.Context, classification models.Classification)) *MocklocalRepo_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Classification))
	})
	return _c
}

func (_c *MocklocalRepo_Put_Call) Return(_a0 error) *MocklocalRepo_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MocklocalRepo_Put_Call) RunAndReturn(run func(context.Context, models.Classification) error) *MocklocalRepo_Put_Call {
	_c.Call.Return(run)
	return _c
}

// NewMocklocalRepo creates a new instance of MocklocalRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMocklocalRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MocklocalRepo {
	mock := &MocklocalRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

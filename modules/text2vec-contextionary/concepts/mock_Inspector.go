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

package concepts

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/weaviate/weaviate/entities/models"
)

// MockInspector is an autogenerated mock type for the Inspector type
type MockInspector struct {
	mock.Mock
}

type MockInspector_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInspector) EXPECT() *MockInspector_Expecter {
	return &MockInspector_Expecter{mock: &_m.Mock}
}

// GetWords provides a mock function with given fields: ctx, words
func (_m *MockInspector) GetWords(ctx context.Context, words string) (*models.C11yWordsResponse, error) {
	ret := _m.Called(ctx, words)

	if len(ret) == 0 {
		panic("no return value specified for GetWords")
	}

	var r0 *models.C11yWordsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.C11yWordsResponse, error)); ok {
		return rf(ctx, words)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.C11yWordsResponse); ok {
		r0 = rf(ctx, words)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.C11yWordsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, words)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInspector_GetWords_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWords'
type MockInspector_GetWords_Call struct {
	*mock.Call
}

// GetWords is a helper method to define mock.On call
//   - ctx context.Context
//   - words string
func (_e *MockInspector_Expecter) GetWords(ctx interface{}, words interface{}) *MockInspector_GetWords_Call {
	return &MockInspector_GetWords_Call{Call: _e.mock.On("GetWords", ctx, words)}
}

func (_c *MockInspector_GetWords_Call) Run(run func(ctx context.Context, words string)) *MockInspector_GetWords_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockInspector_GetWords_Call) Return(_a0 *models.C11yWordsResponse, _a1 error) *MockInspector_GetWords_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInspector_GetWords_Call) RunAndReturn(run func(context.Context, string) (*models.C11yWordsResponse, error)) *MockInspector_GetWords_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInspector creates a new instance of MockInspector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInspector(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInspector {
	mock := &MockInspector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package backups

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	models "github.com/weaviate/weaviate/entities/models"
)

// MockBackupsCreateHandlerFunc is an autogenerated mock type for the BackupsCreateHandlerFunc type
type MockBackupsCreateHandlerFunc struct {
	mock.Mock
}

type MockBackupsCreateHandlerFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBackupsCreateHandlerFunc) EXPECT() *MockBackupsCreateHandlerFunc_Expecter {
	return &MockBackupsCreateHandlerFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockBackupsCreateHandlerFunc) Execute(_a0 BackupsCreateParams, _a1 *models.Principal) middleware.Responder {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(BackupsCreateParams, *models.Principal) middleware.Responder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// MockBackupsCreateHandlerFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockBackupsCreateHandlerFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 BackupsCreateParams
//   - _a1 *models.Principal
func (_e *MockBackupsCreateHandlerFunc_Expecter) Execute(_a0 interface{}, _a1 interface{}) *MockBackupsCreateHandlerFunc_Execute_Call {
	return &MockBackupsCreateHandlerFunc_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *MockBackupsCreateHandlerFunc_Execute_Call) Run(run func(_a0 BackupsCreateParams, _a1 *models.Principal)) *MockBackupsCreateHandlerFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(BackupsCreateParams), args[1].(*models.Principal))
	})
	return _c
}

func (_c *MockBackupsCreateHandlerFunc_Execute_Call) Return(_a0 middleware.Responder) *MockBackupsCreateHandlerFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBackupsCreateHandlerFunc_Execute_Call) RunAndReturn(run func(BackupsCreateParams, *models.Principal) middleware.Responder) *MockBackupsCreateHandlerFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBackupsCreateHandlerFunc creates a new instance of MockBackupsCreateHandlerFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBackupsCreateHandlerFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBackupsCreateHandlerFunc {
	mock := &MockBackupsCreateHandlerFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

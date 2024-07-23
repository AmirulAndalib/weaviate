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

package schema

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	models "github.com/weaviate/weaviate/entities/models"
)

// MockSchemaObjectsPropertiesAddHandler is an autogenerated mock type for the SchemaObjectsPropertiesAddHandler type
type MockSchemaObjectsPropertiesAddHandler struct {
	mock.Mock
}

type MockSchemaObjectsPropertiesAddHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSchemaObjectsPropertiesAddHandler) EXPECT() *MockSchemaObjectsPropertiesAddHandler_Expecter {
	return &MockSchemaObjectsPropertiesAddHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: _a0, _a1
func (_m *MockSchemaObjectsPropertiesAddHandler) Handle(_a0 SchemaObjectsPropertiesAddParams, _a1 *models.Principal) middleware.Responder {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(SchemaObjectsPropertiesAddParams, *models.Principal) middleware.Responder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}

// MockSchemaObjectsPropertiesAddHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockSchemaObjectsPropertiesAddHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - _a0 SchemaObjectsPropertiesAddParams
//   - _a1 *models.Principal
func (_e *MockSchemaObjectsPropertiesAddHandler_Expecter) Handle(_a0 interface{}, _a1 interface{}) *MockSchemaObjectsPropertiesAddHandler_Handle_Call {
	return &MockSchemaObjectsPropertiesAddHandler_Handle_Call{Call: _e.mock.On("Handle", _a0, _a1)}
}

func (_c *MockSchemaObjectsPropertiesAddHandler_Handle_Call) Run(run func(_a0 SchemaObjectsPropertiesAddParams, _a1 *models.Principal)) *MockSchemaObjectsPropertiesAddHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(SchemaObjectsPropertiesAddParams), args[1].(*models.Principal))
	})
	return _c
}

func (_c *MockSchemaObjectsPropertiesAddHandler_Handle_Call) Return(_a0 middleware.Responder) *MockSchemaObjectsPropertiesAddHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSchemaObjectsPropertiesAddHandler_Handle_Call) RunAndReturn(run func(SchemaObjectsPropertiesAddParams, *models.Principal) middleware.Responder) *MockSchemaObjectsPropertiesAddHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSchemaObjectsPropertiesAddHandler creates a new instance of MockSchemaObjectsPropertiesAddHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSchemaObjectsPropertiesAddHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSchemaObjectsPropertiesAddHandler {
	mock := &MockSchemaObjectsPropertiesAddHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package filterext

import (
	mock "github.com/stretchr/testify/mock"
	filters "github.com/weaviate/weaviate/entities/filters"

	models "github.com/weaviate/weaviate/entities/models"
)

// MockvalueExtractorFunc is an autogenerated mock type for the valueExtractorFunc type
type MockvalueExtractorFunc struct {
	mock.Mock
}

type MockvalueExtractorFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockvalueExtractorFunc) EXPECT() *MockvalueExtractorFunc_Expecter {
	return &MockvalueExtractorFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockvalueExtractorFunc) Execute(_a0 *models.WhereFilter) (*filters.Value, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *filters.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.WhereFilter) (*filters.Value, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*models.WhereFilter) *filters.Value); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*filters.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.WhereFilter) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockvalueExtractorFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockvalueExtractorFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *models.WhereFilter
func (_e *MockvalueExtractorFunc_Expecter) Execute(_a0 interface{}) *MockvalueExtractorFunc_Execute_Call {
	return &MockvalueExtractorFunc_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockvalueExtractorFunc_Execute_Call) Run(run func(_a0 *models.WhereFilter)) *MockvalueExtractorFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.WhereFilter))
	})
	return _c
}

func (_c *MockvalueExtractorFunc_Execute_Call) Return(_a0 *filters.Value, _a1 error) *MockvalueExtractorFunc_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockvalueExtractorFunc_Execute_Call) RunAndReturn(run func(*models.WhereFilter) (*filters.Value, error)) *MockvalueExtractorFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockvalueExtractorFunc creates a new instance of MockvalueExtractorFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockvalueExtractorFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockvalueExtractorFunc {
	mock := &MockvalueExtractorFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

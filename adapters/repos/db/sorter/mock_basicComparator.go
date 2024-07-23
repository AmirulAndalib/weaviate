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

import mock "github.com/stretchr/testify/mock"

// MockbasicComparator is an autogenerated mock type for the basicComparator type
type MockbasicComparator struct {
	mock.Mock
}

type MockbasicComparator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockbasicComparator) EXPECT() *MockbasicComparator_Expecter {
	return &MockbasicComparator_Expecter{mock: &_m.Mock}
}

// compare provides a mock function with given fields: a, b
func (_m *MockbasicComparator) compare(a interface{}, b interface{}) int {
	ret := _m.Called(a, b)

	if len(ret) == 0 {
		panic("no return value specified for compare")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) int); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockbasicComparator_compare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'compare'
type MockbasicComparator_compare_Call struct {
	*mock.Call
}

// compare is a helper method to define mock.On call
//   - a interface{}
//   - b interface{}
func (_e *MockbasicComparator_Expecter) compare(a interface{}, b interface{}) *MockbasicComparator_compare_Call {
	return &MockbasicComparator_compare_Call{Call: _e.mock.On("compare", a, b)}
}

func (_c *MockbasicComparator_compare_Call) Run(run func(a interface{}, b interface{})) *MockbasicComparator_compare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(interface{}))
	})
	return _c
}

func (_c *MockbasicComparator_compare_Call) Return(_a0 int) *MockbasicComparator_compare_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockbasicComparator_compare_Call) RunAndReturn(run func(interface{}, interface{}) int) *MockbasicComparator_compare_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockbasicComparator creates a new instance of MockbasicComparator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockbasicComparator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockbasicComparator {
	mock := &MockbasicComparator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

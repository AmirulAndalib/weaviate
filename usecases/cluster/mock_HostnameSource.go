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

package cluster

import mock "github.com/stretchr/testify/mock"

// MockHostnameSource is an autogenerated mock type for the HostnameSource type
type MockHostnameSource struct {
	mock.Mock
}

type MockHostnameSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHostnameSource) EXPECT() *MockHostnameSource_Expecter {
	return &MockHostnameSource_Expecter{mock: &_m.Mock}
}

// AllNames provides a mock function with given fields:
func (_m *MockHostnameSource) AllNames() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AllNames")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockHostnameSource_AllNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllNames'
type MockHostnameSource_AllNames_Call struct {
	*mock.Call
}

// AllNames is a helper method to define mock.On call
func (_e *MockHostnameSource_Expecter) AllNames() *MockHostnameSource_AllNames_Call {
	return &MockHostnameSource_AllNames_Call{Call: _e.mock.On("AllNames")}
}

func (_c *MockHostnameSource_AllNames_Call) Run(run func()) *MockHostnameSource_AllNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHostnameSource_AllNames_Call) Return(_a0 []string) *MockHostnameSource_AllNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHostnameSource_AllNames_Call) RunAndReturn(run func() []string) *MockHostnameSource_AllNames_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHostnameSource creates a new instance of MockHostnameSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHostnameSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHostnameSource {
	mock := &MockHostnameSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

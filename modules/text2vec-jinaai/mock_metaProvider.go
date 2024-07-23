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

package modjinaai

import mock "github.com/stretchr/testify/mock"

// MockmetaProvider is an autogenerated mock type for the metaProvider type
type MockmetaProvider struct {
	mock.Mock
}

type MockmetaProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockmetaProvider) EXPECT() *MockmetaProvider_Expecter {
	return &MockmetaProvider_Expecter{mock: &_m.Mock}
}

// MetaInfo provides a mock function with given fields:
func (_m *MockmetaProvider) MetaInfo() (map[string]interface{}, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MetaInfo")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockmetaProvider_MetaInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetaInfo'
type MockmetaProvider_MetaInfo_Call struct {
	*mock.Call
}

// MetaInfo is a helper method to define mock.On call
func (_e *MockmetaProvider_Expecter) MetaInfo() *MockmetaProvider_MetaInfo_Call {
	return &MockmetaProvider_MetaInfo_Call{Call: _e.mock.On("MetaInfo")}
}

func (_c *MockmetaProvider_MetaInfo_Call) Run(run func()) *MockmetaProvider_MetaInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockmetaProvider_MetaInfo_Call) Return(_a0 map[string]interface{}, _a1 error) *MockmetaProvider_MetaInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockmetaProvider_MetaInfo_Call) RunAndReturn(run func() (map[string]interface{}, error)) *MockmetaProvider_MetaInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockmetaProvider creates a new instance of MockmetaProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockmetaProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockmetaProvider {
	mock := &MockmetaProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package lsmkv

import mock "github.com/stretchr/testify/mock"

// MockBucketOption is an autogenerated mock type for the BucketOption type
type MockBucketOption struct {
	mock.Mock
}

type MockBucketOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBucketOption) EXPECT() *MockBucketOption_Expecter {
	return &MockBucketOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: b
func (_m *MockBucketOption) Execute(b *Bucket) error {
	ret := _m.Called(b)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*Bucket) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBucketOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockBucketOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - b *Bucket
func (_e *MockBucketOption_Expecter) Execute(b interface{}) *MockBucketOption_Execute_Call {
	return &MockBucketOption_Execute_Call{Call: _e.mock.On("Execute", b)}
}

func (_c *MockBucketOption_Execute_Call) Run(run func(b *Bucket)) *MockBucketOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Bucket))
	})
	return _c
}

func (_c *MockBucketOption_Execute_Call) Return(_a0 error) *MockBucketOption_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBucketOption_Execute_Call) RunAndReturn(run func(*Bucket) error) *MockBucketOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBucketOption creates a new instance of MockBucketOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBucketOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBucketOption {
	mock := &MockBucketOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

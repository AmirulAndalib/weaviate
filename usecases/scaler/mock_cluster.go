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

package scaler

import mock "github.com/stretchr/testify/mock"

// Mockcluster is an autogenerated mock type for the cluster type
type Mockcluster struct {
	mock.Mock
}

type Mockcluster_Expecter struct {
	mock *mock.Mock
}

func (_m *Mockcluster) EXPECT() *Mockcluster_Expecter {
	return &Mockcluster_Expecter{mock: &_m.Mock}
}

// Candidates provides a mock function with given fields:
func (_m *Mockcluster) Candidates() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Candidates")
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

// Mockcluster_Candidates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Candidates'
type Mockcluster_Candidates_Call struct {
	*mock.Call
}

// Candidates is a helper method to define mock.On call
func (_e *Mockcluster_Expecter) Candidates() *Mockcluster_Candidates_Call {
	return &Mockcluster_Candidates_Call{Call: _e.mock.On("Candidates")}
}

func (_c *Mockcluster_Candidates_Call) Run(run func()) *Mockcluster_Candidates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Mockcluster_Candidates_Call) Return(_a0 []string) *Mockcluster_Candidates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockcluster_Candidates_Call) RunAndReturn(run func() []string) *Mockcluster_Candidates_Call {
	_c.Call.Return(run)
	return _c
}

// LocalName provides a mock function with given fields:
func (_m *Mockcluster) LocalName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LocalName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Mockcluster_LocalName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LocalName'
type Mockcluster_LocalName_Call struct {
	*mock.Call
}

// LocalName is a helper method to define mock.On call
func (_e *Mockcluster_Expecter) LocalName() *Mockcluster_LocalName_Call {
	return &Mockcluster_LocalName_Call{Call: _e.mock.On("LocalName")}
}

func (_c *Mockcluster_LocalName_Call) Run(run func()) *Mockcluster_LocalName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Mockcluster_LocalName_Call) Return(_a0 string) *Mockcluster_LocalName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockcluster_LocalName_Call) RunAndReturn(run func() string) *Mockcluster_LocalName_Call {
	_c.Call.Return(run)
	return _c
}

// NodeHostname provides a mock function with given fields: name
func (_m *Mockcluster) NodeHostname(name string) (string, bool) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for NodeHostname")
	}

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (string, bool)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Mockcluster_NodeHostname_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NodeHostname'
type Mockcluster_NodeHostname_Call struct {
	*mock.Call
}

// NodeHostname is a helper method to define mock.On call
//   - name string
func (_e *Mockcluster_Expecter) NodeHostname(name interface{}) *Mockcluster_NodeHostname_Call {
	return &Mockcluster_NodeHostname_Call{Call: _e.mock.On("NodeHostname", name)}
}

func (_c *Mockcluster_NodeHostname_Call) Run(run func(name string)) *Mockcluster_NodeHostname_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Mockcluster_NodeHostname_Call) Return(_a0 string, _a1 bool) *Mockcluster_NodeHostname_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Mockcluster_NodeHostname_Call) RunAndReturn(run func(string) (string, bool)) *Mockcluster_NodeHostname_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockcluster creates a new instance of Mockcluster. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockcluster(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mockcluster {
	mock := &Mockcluster{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

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

package vectorizer

import mock "github.com/stretchr/testify/mock"

// MockClassSettings is an autogenerated mock type for the ClassSettings type
type MockClassSettings struct {
	mock.Mock
}

type MockClassSettings_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClassSettings) EXPECT() *MockClassSettings_Expecter {
	return &MockClassSettings_Expecter{mock: &_m.Mock}
}

// EndpointURL provides a mock function with given fields:
func (_m *MockClassSettings) EndpointURL() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EndpointURL")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClassSettings_EndpointURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndpointURL'
type MockClassSettings_EndpointURL_Call struct {
	*mock.Call
}

// EndpointURL is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) EndpointURL() *MockClassSettings_EndpointURL_Call {
	return &MockClassSettings_EndpointURL_Call{Call: _e.mock.On("EndpointURL")}
}

func (_c *MockClassSettings_EndpointURL_Call) Run(run func()) *MockClassSettings_EndpointURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_EndpointURL_Call) Return(_a0 string) *MockClassSettings_EndpointURL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_EndpointURL_Call) RunAndReturn(run func() string) *MockClassSettings_EndpointURL_Call {
	_c.Call.Return(run)
	return _c
}

// OptionUseCache provides a mock function with given fields:
func (_m *MockClassSettings) OptionUseCache() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OptionUseCache")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClassSettings_OptionUseCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionUseCache'
type MockClassSettings_OptionUseCache_Call struct {
	*mock.Call
}

// OptionUseCache is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) OptionUseCache() *MockClassSettings_OptionUseCache_Call {
	return &MockClassSettings_OptionUseCache_Call{Call: _e.mock.On("OptionUseCache")}
}

func (_c *MockClassSettings_OptionUseCache_Call) Run(run func()) *MockClassSettings_OptionUseCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_OptionUseCache_Call) Return(_a0 bool) *MockClassSettings_OptionUseCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_OptionUseCache_Call) RunAndReturn(run func() bool) *MockClassSettings_OptionUseCache_Call {
	_c.Call.Return(run)
	return _c
}

// OptionUseGPU provides a mock function with given fields:
func (_m *MockClassSettings) OptionUseGPU() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OptionUseGPU")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClassSettings_OptionUseGPU_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionUseGPU'
type MockClassSettings_OptionUseGPU_Call struct {
	*mock.Call
}

// OptionUseGPU is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) OptionUseGPU() *MockClassSettings_OptionUseGPU_Call {
	return &MockClassSettings_OptionUseGPU_Call{Call: _e.mock.On("OptionUseGPU")}
}

func (_c *MockClassSettings_OptionUseGPU_Call) Run(run func()) *MockClassSettings_OptionUseGPU_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_OptionUseGPU_Call) Return(_a0 bool) *MockClassSettings_OptionUseGPU_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_OptionUseGPU_Call) RunAndReturn(run func() bool) *MockClassSettings_OptionUseGPU_Call {
	_c.Call.Return(run)
	return _c
}

// OptionWaitForModel provides a mock function with given fields:
func (_m *MockClassSettings) OptionWaitForModel() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OptionWaitForModel")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClassSettings_OptionWaitForModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionWaitForModel'
type MockClassSettings_OptionWaitForModel_Call struct {
	*mock.Call
}

// OptionWaitForModel is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) OptionWaitForModel() *MockClassSettings_OptionWaitForModel_Call {
	return &MockClassSettings_OptionWaitForModel_Call{Call: _e.mock.On("OptionWaitForModel")}
}

func (_c *MockClassSettings_OptionWaitForModel_Call) Run(run func()) *MockClassSettings_OptionWaitForModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_OptionWaitForModel_Call) Return(_a0 bool) *MockClassSettings_OptionWaitForModel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_OptionWaitForModel_Call) RunAndReturn(run func() bool) *MockClassSettings_OptionWaitForModel_Call {
	_c.Call.Return(run)
	return _c
}

// PassageModel provides a mock function with given fields:
func (_m *MockClassSettings) PassageModel() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PassageModel")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClassSettings_PassageModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PassageModel'
type MockClassSettings_PassageModel_Call struct {
	*mock.Call
}

// PassageModel is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) PassageModel() *MockClassSettings_PassageModel_Call {
	return &MockClassSettings_PassageModel_Call{Call: _e.mock.On("PassageModel")}
}

func (_c *MockClassSettings_PassageModel_Call) Run(run func()) *MockClassSettings_PassageModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_PassageModel_Call) Return(_a0 string) *MockClassSettings_PassageModel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_PassageModel_Call) RunAndReturn(run func() string) *MockClassSettings_PassageModel_Call {
	_c.Call.Return(run)
	return _c
}

// PropertyIndexed provides a mock function with given fields: property
func (_m *MockClassSettings) PropertyIndexed(property string) bool {
	ret := _m.Called(property)

	if len(ret) == 0 {
		panic("no return value specified for PropertyIndexed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(property)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClassSettings_PropertyIndexed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PropertyIndexed'
type MockClassSettings_PropertyIndexed_Call struct {
	*mock.Call
}

// PropertyIndexed is a helper method to define mock.On call
//   - property string
func (_e *MockClassSettings_Expecter) PropertyIndexed(property interface{}) *MockClassSettings_PropertyIndexed_Call {
	return &MockClassSettings_PropertyIndexed_Call{Call: _e.mock.On("PropertyIndexed", property)}
}

func (_c *MockClassSettings_PropertyIndexed_Call) Run(run func(property string)) *MockClassSettings_PropertyIndexed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClassSettings_PropertyIndexed_Call) Return(_a0 bool) *MockClassSettings_PropertyIndexed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_PropertyIndexed_Call) RunAndReturn(run func(string) bool) *MockClassSettings_PropertyIndexed_Call {
	_c.Call.Return(run)
	return _c
}

// QueryModel provides a mock function with given fields:
func (_m *MockClassSettings) QueryModel() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for QueryModel")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClassSettings_QueryModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryModel'
type MockClassSettings_QueryModel_Call struct {
	*mock.Call
}

// QueryModel is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) QueryModel() *MockClassSettings_QueryModel_Call {
	return &MockClassSettings_QueryModel_Call{Call: _e.mock.On("QueryModel")}
}

func (_c *MockClassSettings_QueryModel_Call) Run(run func()) *MockClassSettings_QueryModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_QueryModel_Call) Return(_a0 string) *MockClassSettings_QueryModel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_QueryModel_Call) RunAndReturn(run func() string) *MockClassSettings_QueryModel_Call {
	_c.Call.Return(run)
	return _c
}

// VectorizeClassName provides a mock function with given fields:
func (_m *MockClassSettings) VectorizeClassName() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for VectorizeClassName")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClassSettings_VectorizeClassName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VectorizeClassName'
type MockClassSettings_VectorizeClassName_Call struct {
	*mock.Call
}

// VectorizeClassName is a helper method to define mock.On call
func (_e *MockClassSettings_Expecter) VectorizeClassName() *MockClassSettings_VectorizeClassName_Call {
	return &MockClassSettings_VectorizeClassName_Call{Call: _e.mock.On("VectorizeClassName")}
}

func (_c *MockClassSettings_VectorizeClassName_Call) Run(run func()) *MockClassSettings_VectorizeClassName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClassSettings_VectorizeClassName_Call) Return(_a0 bool) *MockClassSettings_VectorizeClassName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_VectorizeClassName_Call) RunAndReturn(run func() bool) *MockClassSettings_VectorizeClassName_Call {
	_c.Call.Return(run)
	return _c
}

// VectorizePropertyName provides a mock function with given fields: propertyName
func (_m *MockClassSettings) VectorizePropertyName(propertyName string) bool {
	ret := _m.Called(propertyName)

	if len(ret) == 0 {
		panic("no return value specified for VectorizePropertyName")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(propertyName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClassSettings_VectorizePropertyName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VectorizePropertyName'
type MockClassSettings_VectorizePropertyName_Call struct {
	*mock.Call
}

// VectorizePropertyName is a helper method to define mock.On call
//   - propertyName string
func (_e *MockClassSettings_Expecter) VectorizePropertyName(propertyName interface{}) *MockClassSettings_VectorizePropertyName_Call {
	return &MockClassSettings_VectorizePropertyName_Call{Call: _e.mock.On("VectorizePropertyName", propertyName)}
}

func (_c *MockClassSettings_VectorizePropertyName_Call) Run(run func(propertyName string)) *MockClassSettings_VectorizePropertyName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClassSettings_VectorizePropertyName_Call) Return(_a0 bool) *MockClassSettings_VectorizePropertyName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClassSettings_VectorizePropertyName_Call) RunAndReturn(run func(string) bool) *MockClassSettings_VectorizePropertyName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClassSettings creates a new instance of MockClassSettings. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClassSettings(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClassSettings {
	mock := &MockClassSettings{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

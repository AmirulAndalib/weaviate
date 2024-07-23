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

package modgenerativedummy

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	moduletools "github.com/weaviate/weaviate/entities/moduletools"
	models "github.com/weaviate/weaviate/usecases/modulecomponents/additional/models"
)

// MockgenerativeClient is an autogenerated mock type for the generativeClient type
type MockgenerativeClient struct {
	mock.Mock
}

type MockgenerativeClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockgenerativeClient) EXPECT() *MockgenerativeClient_Expecter {
	return &MockgenerativeClient_Expecter{mock: &_m.Mock}
}

// Generate provides a mock function with given fields: ctx, cfg, prompt
func (_m *MockgenerativeClient) Generate(ctx context.Context, cfg moduletools.ClassConfig, prompt string) (*models.GenerateResponse, error) {
	ret := _m.Called(ctx, cfg, prompt)

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 *models.GenerateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, moduletools.ClassConfig, string) (*models.GenerateResponse, error)); ok {
		return rf(ctx, cfg, prompt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, moduletools.ClassConfig, string) *models.GenerateResponse); ok {
		r0 = rf(ctx, cfg, prompt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GenerateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, moduletools.ClassConfig, string) error); ok {
		r1 = rf(ctx, cfg, prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockgenerativeClient_Generate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Generate'
type MockgenerativeClient_Generate_Call struct {
	*mock.Call
}

// Generate is a helper method to define mock.On call
//   - ctx context.Context
//   - cfg moduletools.ClassConfig
//   - prompt string
func (_e *MockgenerativeClient_Expecter) Generate(ctx interface{}, cfg interface{}, prompt interface{}) *MockgenerativeClient_Generate_Call {
	return &MockgenerativeClient_Generate_Call{Call: _e.mock.On("Generate", ctx, cfg, prompt)}
}

func (_c *MockgenerativeClient_Generate_Call) Run(run func(ctx context.Context, cfg moduletools.ClassConfig, prompt string)) *MockgenerativeClient_Generate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(moduletools.ClassConfig), args[2].(string))
	})
	return _c
}

func (_c *MockgenerativeClient_Generate_Call) Return(_a0 *models.GenerateResponse, _a1 error) *MockgenerativeClient_Generate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockgenerativeClient_Generate_Call) RunAndReturn(run func(context.Context, moduletools.ClassConfig, string) (*models.GenerateResponse, error)) *MockgenerativeClient_Generate_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateAllResults provides a mock function with given fields: ctx, textProperties, task, cfg
func (_m *MockgenerativeClient) GenerateAllResults(ctx context.Context, textProperties []map[string]string, task string, cfg moduletools.ClassConfig) (*models.GenerateResponse, error) {
	ret := _m.Called(ctx, textProperties, task, cfg)

	if len(ret) == 0 {
		panic("no return value specified for GenerateAllResults")
	}

	var r0 *models.GenerateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []map[string]string, string, moduletools.ClassConfig) (*models.GenerateResponse, error)); ok {
		return rf(ctx, textProperties, task, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []map[string]string, string, moduletools.ClassConfig) *models.GenerateResponse); ok {
		r0 = rf(ctx, textProperties, task, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GenerateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []map[string]string, string, moduletools.ClassConfig) error); ok {
		r1 = rf(ctx, textProperties, task, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockgenerativeClient_GenerateAllResults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateAllResults'
type MockgenerativeClient_GenerateAllResults_Call struct {
	*mock.Call
}

// GenerateAllResults is a helper method to define mock.On call
//   - ctx context.Context
//   - textProperties []map[string]string
//   - task string
//   - cfg moduletools.ClassConfig
func (_e *MockgenerativeClient_Expecter) GenerateAllResults(ctx interface{}, textProperties interface{}, task interface{}, cfg interface{}) *MockgenerativeClient_GenerateAllResults_Call {
	return &MockgenerativeClient_GenerateAllResults_Call{Call: _e.mock.On("GenerateAllResults", ctx, textProperties, task, cfg)}
}

func (_c *MockgenerativeClient_GenerateAllResults_Call) Run(run func(ctx context.Context, textProperties []map[string]string, task string, cfg moduletools.ClassConfig)) *MockgenerativeClient_GenerateAllResults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]map[string]string), args[2].(string), args[3].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MockgenerativeClient_GenerateAllResults_Call) Return(_a0 *models.GenerateResponse, _a1 error) *MockgenerativeClient_GenerateAllResults_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockgenerativeClient_GenerateAllResults_Call) RunAndReturn(run func(context.Context, []map[string]string, string, moduletools.ClassConfig) (*models.GenerateResponse, error)) *MockgenerativeClient_GenerateAllResults_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateSingleResult provides a mock function with given fields: ctx, textProperties, prompt, cfg
func (_m *MockgenerativeClient) GenerateSingleResult(ctx context.Context, textProperties map[string]string, prompt string, cfg moduletools.ClassConfig) (*models.GenerateResponse, error) {
	ret := _m.Called(ctx, textProperties, prompt, cfg)

	if len(ret) == 0 {
		panic("no return value specified for GenerateSingleResult")
	}

	var r0 *models.GenerateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, string, moduletools.ClassConfig) (*models.GenerateResponse, error)); ok {
		return rf(ctx, textProperties, prompt, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, string, moduletools.ClassConfig) *models.GenerateResponse); ok {
		r0 = rf(ctx, textProperties, prompt, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GenerateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]string, string, moduletools.ClassConfig) error); ok {
		r1 = rf(ctx, textProperties, prompt, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockgenerativeClient_GenerateSingleResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateSingleResult'
type MockgenerativeClient_GenerateSingleResult_Call struct {
	*mock.Call
}

// GenerateSingleResult is a helper method to define mock.On call
//   - ctx context.Context
//   - textProperties map[string]string
//   - prompt string
//   - cfg moduletools.ClassConfig
func (_e *MockgenerativeClient_Expecter) GenerateSingleResult(ctx interface{}, textProperties interface{}, prompt interface{}, cfg interface{}) *MockgenerativeClient_GenerateSingleResult_Call {
	return &MockgenerativeClient_GenerateSingleResult_Call{Call: _e.mock.On("GenerateSingleResult", ctx, textProperties, prompt, cfg)}
}

func (_c *MockgenerativeClient_GenerateSingleResult_Call) Run(run func(ctx context.Context, textProperties map[string]string, prompt string, cfg moduletools.ClassConfig)) *MockgenerativeClient_GenerateSingleResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]string), args[2].(string), args[3].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MockgenerativeClient_GenerateSingleResult_Call) Return(_a0 *models.GenerateResponse, _a1 error) *MockgenerativeClient_GenerateSingleResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockgenerativeClient_GenerateSingleResult_Call) RunAndReturn(run func(context.Context, map[string]string, string, moduletools.ClassConfig) (*models.GenerateResponse, error)) *MockgenerativeClient_GenerateSingleResult_Call {
	_c.Call.Return(run)
	return _c
}

// MetaInfo provides a mock function with given fields:
func (_m *MockgenerativeClient) MetaInfo() (map[string]interface{}, error) {
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

// MockgenerativeClient_MetaInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetaInfo'
type MockgenerativeClient_MetaInfo_Call struct {
	*mock.Call
}

// MetaInfo is a helper method to define mock.On call
func (_e *MockgenerativeClient_Expecter) MetaInfo() *MockgenerativeClient_MetaInfo_Call {
	return &MockgenerativeClient_MetaInfo_Call{Call: _e.mock.On("MetaInfo")}
}

func (_c *MockgenerativeClient_MetaInfo_Call) Run(run func()) *MockgenerativeClient_MetaInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockgenerativeClient_MetaInfo_Call) Return(_a0 map[string]interface{}, _a1 error) *MockgenerativeClient_MetaInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockgenerativeClient_MetaInfo_Call) RunAndReturn(run func() (map[string]interface{}, error)) *MockgenerativeClient_MetaInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockgenerativeClient creates a new instance of MockgenerativeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockgenerativeClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockgenerativeClient {
	mock := &MockgenerativeClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

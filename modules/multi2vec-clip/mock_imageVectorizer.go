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

package modclip

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/weaviate/weaviate/entities/models"

	moduletools "github.com/weaviate/weaviate/entities/moduletools"
)

// MockimageVectorizer is an autogenerated mock type for the imageVectorizer type
type MockimageVectorizer struct {
	mock.Mock
}

type MockimageVectorizer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockimageVectorizer) EXPECT() *MockimageVectorizer_Expecter {
	return &MockimageVectorizer_Expecter{mock: &_m.Mock}
}

// Object provides a mock function with given fields: ctx, obj, comp, cfg
func (_m *MockimageVectorizer) Object(ctx context.Context, obj *models.Object, comp moduletools.VectorizablePropsComparator, cfg moduletools.ClassConfig) ([]float32, models.AdditionalProperties, error) {
	ret := _m.Called(ctx, obj, comp, cfg)

	if len(ret) == 0 {
		panic("no return value specified for Object")
	}

	var r0 []float32
	var r1 models.AdditionalProperties
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Object, moduletools.VectorizablePropsComparator, moduletools.ClassConfig) ([]float32, models.AdditionalProperties, error)); ok {
		return rf(ctx, obj, comp, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Object, moduletools.VectorizablePropsComparator, moduletools.ClassConfig) []float32); ok {
		r0 = rf(ctx, obj, comp, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Object, moduletools.VectorizablePropsComparator, moduletools.ClassConfig) models.AdditionalProperties); ok {
		r1 = rf(ctx, obj, comp, cfg)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(models.AdditionalProperties)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *models.Object, moduletools.VectorizablePropsComparator, moduletools.ClassConfig) error); ok {
		r2 = rf(ctx, obj, comp, cfg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockimageVectorizer_Object_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Object'
type MockimageVectorizer_Object_Call struct {
	*mock.Call
}

// Object is a helper method to define mock.On call
//   - ctx context.Context
//   - obj *models.Object
//   - comp moduletools.VectorizablePropsComparator
//   - cfg moduletools.ClassConfig
func (_e *MockimageVectorizer_Expecter) Object(ctx interface{}, obj interface{}, comp interface{}, cfg interface{}) *MockimageVectorizer_Object_Call {
	return &MockimageVectorizer_Object_Call{Call: _e.mock.On("Object", ctx, obj, comp, cfg)}
}

func (_c *MockimageVectorizer_Object_Call) Run(run func(ctx context.Context, obj *models.Object, comp moduletools.VectorizablePropsComparator, cfg moduletools.ClassConfig)) *MockimageVectorizer_Object_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Object), args[2].(moduletools.VectorizablePropsComparator), args[3].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MockimageVectorizer_Object_Call) Return(_a0 []float32, _a1 models.AdditionalProperties, _a2 error) *MockimageVectorizer_Object_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockimageVectorizer_Object_Call) RunAndReturn(run func(context.Context, *models.Object, moduletools.VectorizablePropsComparator, moduletools.ClassConfig) ([]float32, models.AdditionalProperties, error)) *MockimageVectorizer_Object_Call {
	_c.Call.Return(run)
	return _c
}

// VectorizeImage provides a mock function with given fields: ctx, id, image, cfg
func (_m *MockimageVectorizer) VectorizeImage(ctx context.Context, id string, image string, cfg moduletools.ClassConfig) ([]float32, error) {
	ret := _m.Called(ctx, id, image, cfg)

	if len(ret) == 0 {
		panic("no return value specified for VectorizeImage")
	}

	var r0 []float32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, moduletools.ClassConfig) ([]float32, error)); ok {
		return rf(ctx, id, image, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, moduletools.ClassConfig) []float32); ok {
		r0 = rf(ctx, id, image, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, moduletools.ClassConfig) error); ok {
		r1 = rf(ctx, id, image, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockimageVectorizer_VectorizeImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VectorizeImage'
type MockimageVectorizer_VectorizeImage_Call struct {
	*mock.Call
}

// VectorizeImage is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - image string
//   - cfg moduletools.ClassConfig
func (_e *MockimageVectorizer_Expecter) VectorizeImage(ctx interface{}, id interface{}, image interface{}, cfg interface{}) *MockimageVectorizer_VectorizeImage_Call {
	return &MockimageVectorizer_VectorizeImage_Call{Call: _e.mock.On("VectorizeImage", ctx, id, image, cfg)}
}

func (_c *MockimageVectorizer_VectorizeImage_Call) Run(run func(ctx context.Context, id string, image string, cfg moduletools.ClassConfig)) *MockimageVectorizer_VectorizeImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(moduletools.ClassConfig))
	})
	return _c
}

func (_c *MockimageVectorizer_VectorizeImage_Call) Return(_a0 []float32, _a1 error) *MockimageVectorizer_VectorizeImage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockimageVectorizer_VectorizeImage_Call) RunAndReturn(run func(context.Context, string, string, moduletools.ClassConfig) ([]float32, error)) *MockimageVectorizer_VectorizeImage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockimageVectorizer creates a new instance of MockimageVectorizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockimageVectorizer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockimageVectorizer {
	mock := &MockimageVectorizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

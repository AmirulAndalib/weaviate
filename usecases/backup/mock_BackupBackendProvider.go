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

package backup

import (
	mock "github.com/stretchr/testify/mock"
	modulecapabilities "github.com/weaviate/weaviate/entities/modulecapabilities"
)

// MockBackupBackendProvider is an autogenerated mock type for the BackupBackendProvider type
type MockBackupBackendProvider struct {
	mock.Mock
}

type MockBackupBackendProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBackupBackendProvider) EXPECT() *MockBackupBackendProvider_Expecter {
	return &MockBackupBackendProvider_Expecter{mock: &_m.Mock}
}

// BackupBackend provides a mock function with given fields: backend
func (_m *MockBackupBackendProvider) BackupBackend(backend string) (modulecapabilities.BackupBackend, error) {
	ret := _m.Called(backend)

	if len(ret) == 0 {
		panic("no return value specified for BackupBackend")
	}

	var r0 modulecapabilities.BackupBackend
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (modulecapabilities.BackupBackend, error)); ok {
		return rf(backend)
	}
	if rf, ok := ret.Get(0).(func(string) modulecapabilities.BackupBackend); ok {
		r0 = rf(backend)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(modulecapabilities.BackupBackend)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(backend)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBackupBackendProvider_BackupBackend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BackupBackend'
type MockBackupBackendProvider_BackupBackend_Call struct {
	*mock.Call
}

// BackupBackend is a helper method to define mock.On call
//   - backend string
func (_e *MockBackupBackendProvider_Expecter) BackupBackend(backend interface{}) *MockBackupBackendProvider_BackupBackend_Call {
	return &MockBackupBackendProvider_BackupBackend_Call{Call: _e.mock.On("BackupBackend", backend)}
}

func (_c *MockBackupBackendProvider_BackupBackend_Call) Run(run func(backend string)) *MockBackupBackendProvider_BackupBackend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBackupBackendProvider_BackupBackend_Call) Return(_a0 modulecapabilities.BackupBackend, _a1 error) *MockBackupBackendProvider_BackupBackend_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBackupBackendProvider_BackupBackend_Call) RunAndReturn(run func(string) (modulecapabilities.BackupBackend, error)) *MockBackupBackendProvider_BackupBackend_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBackupBackendProvider creates a new instance of MockBackupBackendProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBackupBackendProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBackupBackendProvider {
	mock := &MockBackupBackendProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.14.0. DO NOT EDIT.

package database

import mock "github.com/stretchr/testify/mock"

// MockITransaction is an autogenerated mock type for the ITransaction type
type MockITransaction struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *MockITransaction) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: sqlQuery, params
func (_m *MockITransaction) Exec(sqlQuery string, params ...interface{}) error {
	ret := _m.Called(sqlQuery, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(sqlQuery, params...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *MockITransaction) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Select provides a mock function with given fields: dest, sqlQuery, params
func (_m *MockITransaction) Select(dest interface{}, sqlQuery string, params ...interface{}) error {
	ret := _m.Called(dest, sqlQuery, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, sqlQuery, params...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockITransaction interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockITransaction creates a new instance of MockITransaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockITransaction(t mockConstructorTestingTNewMockITransaction) *MockITransaction {
	mock := &MockITransaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

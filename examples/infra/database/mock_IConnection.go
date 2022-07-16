// Code generated by mockery v2.14.0. DO NOT EDIT.

package database

import mock "github.com/stretchr/testify/mock"

// MockIConnection is an autogenerated mock type for the IConnection type
type MockIConnection struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *MockIConnection) Begin() (ITransaction, error) {
	ret := _m.Called()

	var r0 ITransaction
	if rf, ok := ret.Get(0).(func() ITransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ITransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exec provides a mock function with given fields: sqlQuery, params
func (_m *MockIConnection) Exec(sqlQuery string, params ...interface{}) error {
	ret := _m.Called(sqlQuery, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(sqlQuery, params...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Select provides a mock function with given fields: dest, sqlQuery, params
func (_m *MockIConnection) Select(dest interface{}, sqlQuery string, params ...interface{}) error {
	ret := _m.Called(dest, sqlQuery, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, sqlQuery, params...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockIConnection interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIConnection creates a new instance of MockIConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIConnection(t mockConstructorTestingTNewMockIConnection) *MockIConnection {
	mock := &MockIConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
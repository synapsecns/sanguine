// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Chain is an autogenerated mock type for the Chain type
type Chain struct {
	mock.Mock
}

// ConfirmationsThreshold provides a mock function with given fields:
func (_m *Chain) ConfirmationsThreshold() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Chain) ID() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// URLs provides a mock function with given fields:
func (_m *Chain) URLs() []string {
	ret := _m.Called()

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

type mockConstructorTestingTNewChain interface {
	mock.TestingT
	Cleanup(func())
}

// NewChain creates a new instance of Chain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChain(t mockConstructorTestingTNewChain) *Chain {
	mock := &Chain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

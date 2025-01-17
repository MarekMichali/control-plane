// Code generated by mockery v2.14.0. DO NOT EDIT.

package automock

import (
	ias "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/ias"
	mock "github.com/stretchr/testify/mock"
)

// BundleBuilder is an autogenerated mock type for the BundleBuilder type
type BundleBuilder struct {
	mock.Mock
}

// NewBundle provides a mock function with given fields: identifier, inputID
func (_m *BundleBuilder) NewBundle(identifier string, inputID ias.SPInputID) (ias.Bundle, error) {
	ret := _m.Called(identifier, inputID)

	var r0 ias.Bundle
	if rf, ok := ret.Get(0).(func(string, ias.SPInputID) ias.Bundle); ok {
		r0 = rf(identifier, inputID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ias.Bundle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ias.SPInputID) error); ok {
		r1 = rf(identifier, inputID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBundleBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewBundleBuilder creates a new instance of BundleBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBundleBuilder(t mockConstructorTestingTNewBundleBuilder) *BundleBuilder {
	mock := &BundleBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

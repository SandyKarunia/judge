// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// System is an autogenerated mock type for the System type
type System struct {
	mock.Mock
}

// IsSudo provides a mock function with given fields:
func (_m *System) IsSudo() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

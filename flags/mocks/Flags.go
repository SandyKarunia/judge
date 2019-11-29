// Code generated by mockery v1.0.0
package mocks

import (
	flags "github.com/sandykarunia/fudge/flags"
	mock "github.com/stretchr/testify/mock"
)

// Flags is an autogenerated mock type for the Flags type
type Flags struct {
	mock.Mock
}

// GetBool provides a mock function with given fields: flagName
func (_m *Flags) GetBool(flagName flags.FlagName) bool {
	ret := _m.Called(flagName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flags.FlagName) bool); ok {
		r0 = rf(flagName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

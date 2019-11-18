// Code generated by mockery v1.0.0
package mocks

import (
	language "github.com/sandykarunia/fudge/language"
	mock "github.com/stretchr/testify/mock"
)

// System is an autogenerated mock type for the System type
type System struct {
	mock.Mock
}

// Execute provides a mock function with given fields: cmd, args
func (_m *System) Execute(cmd string, args ...string) (string, error) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, cmd)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(cmd, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(cmd, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHMACSecretFromEnv provides a mock function with given fields:
func (_m *System) GetHMACSecretFromEnv() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsControlGroupSupported provides a mock function with given fields:
func (_m *System) IsControlGroupSupported() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsLanguageSupported provides a mock function with given fields: lang
func (_m *System) IsLanguageSupported(lang language.Language) bool {
	ret := _m.Called(lang)

	var r0 bool
	if rf, ok := ret.Get(0).(func(language.Language) bool); ok {
		r0 = rf(lang)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
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

// VerifyPkgInstalled provides a mock function with given fields: pkgName
func (_m *System) VerifyPkgInstalled(pkgName string) error {
	ret := _m.Called(pkgName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(pkgName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

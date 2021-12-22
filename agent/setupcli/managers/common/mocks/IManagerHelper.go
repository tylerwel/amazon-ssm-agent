// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IManagerHelper is an autogenerated mock type for the IManagerHelper type
type IManagerHelper struct {
	mock.Mock
}

// GetExitCode provides a mock function with given fields: err
func (_m *IManagerHelper) GetExitCode(err error) int {
	ret := _m.Called(err)

	var r0 int
	if rf, ok := ret.Get(0).(func(error) int); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IsCommandAvailable provides a mock function with given fields: cmd
func (_m *IManagerHelper) IsCommandAvailable(cmd string) bool {
	ret := _m.Called(cmd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(cmd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsExitCodeError provides a mock function with given fields: err
func (_m *IManagerHelper) IsExitCodeError(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsTimeoutError provides a mock function with given fields: err
func (_m *IManagerHelper) IsTimeoutError(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RunCommand provides a mock function with given fields: cmd, args
func (_m *IManagerHelper) RunCommand(cmd string, args ...string) (string, error) {
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
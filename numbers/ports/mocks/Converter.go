// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Converter is an autogenerated mock type for the Converter type
type Converter struct {
	mock.Mock
}

// IsSupported provides a mock function with given fields: lang
func (_m *Converter) IsSupported(lang string) error {
	ret := _m.Called(lang)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(lang)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// To provides a mock function with given fields: number, lang
func (_m *Converter) To(number int, lang string) string {
	ret := _m.Called(number, lang)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, string) string); ok {
		r0 = rf(number, lang)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

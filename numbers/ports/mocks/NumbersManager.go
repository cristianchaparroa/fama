// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NumbersManager is an autogenerated mock type for the NumbersManager type
type NumbersManager struct {
	mock.Mock
}

// ToWords provides a mock function with given fields: number, lang
func (_m *NumbersManager) ToWords(number int, lang string) (string, error) {
	ret := _m.Called(number, lang)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, string) string); ok {
		r0 = rf(number, lang)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(number, lang)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/ciazhar/golang-grpc/server-grpc/pkg/social/model"
import rest "github.com/ciazhar/golang-grpc/common/rest"

// SocialUseCase is an autogenerated mock type for the SocialUseCase type
type SocialUseCase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *SocialUseCase) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: param
func (_m *SocialUseCase) Fetch(param rest.Param) ([]model.Social, error) {
	ret := _m.Called(param)

	var r0 []model.Social
	if rf, ok := ret.Get(0).(func(rest.Param) []model.Social); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Social)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(rest.Param) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *SocialUseCase) GetByID(id string) (model.Social, error) {
	ret := _m.Called(id)

	var r0 model.Social
	if rf, ok := ret.Get(0).(func(string) model.Social); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Social)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: req
func (_m *SocialUseCase) Store(req *model.Social) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Social) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: req
func (_m *SocialUseCase) Update(req *model.Social) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Social) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
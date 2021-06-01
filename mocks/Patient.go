// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	model "phudt/week4/internal/model"

	mock "github.com/stretchr/testify/mock"

	repo "phudt/week4/internal/repo"
)

// Patient is an autogenerated mock type for the Patient type
type Patient struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Patient) Create(_a0 model.Patient) (*model.Patient, error) {
	ret := _m.Called(_a0)

	var r0 *model.Patient
	if rf, ok := ret.Get(0).(func(model.Patient) *model.Patient); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Patient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Patient) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: q, filters, sort, page
func (_m *Patient) List(q string, filters []repo.Filter, sort repo.Sort, page repo.Pagination) ([]model.Patient, error) {
	ret := _m.Called(q, filters, sort, page)

	var r0 []model.Patient
	if rf, ok := ret.Get(0).(func(string, []repo.Filter, repo.Sort, repo.Pagination) []model.Patient); ok {
		r0 = rf(q, filters, sort, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Patient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []repo.Filter, repo.Sort, repo.Pagination) error); ok {
		r1 = rf(q, filters, sort, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

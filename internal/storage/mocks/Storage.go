// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	models "Gl0ven/kata_projects/rates/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// SaveRates provides a mock function with given fields: ctx, rates
func (_m *Storage) SaveRates(ctx context.Context, rates models.Rates) error {
	ret := _m.Called(ctx, rates)

	if len(ret) == 0 {
		panic("no return value specified for SaveRates")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Rates) error); ok {
		r0 = rf(ctx, rates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/NodeFactoryIo/vedran/internal/models"

// MetricsRepository is an autogenerated mock type for the MetricsRepository type
type MetricsRepository struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: ID
func (_m *MetricsRepository) FindByID(ID string) (*models.Metrics, error) {
	ret := _m.Called(ID)

	var r0 *models.Metrics
	if rf, ok := ret.Get(0).(func(string) *models.Metrics); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Metrics)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *MetricsRepository) GetAll() (*[]models.Metrics, error) {
	ret := _m.Called()

	var r0 *[]models.Metrics
	if rf, ok := ret.Get(0).(func() *[]models.Metrics); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Metrics)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: metrics
func (_m *MetricsRepository) Save(metrics *models.Metrics) error {
	ret := _m.Called(metrics)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Metrics) error); ok {
		r0 = rf(metrics)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

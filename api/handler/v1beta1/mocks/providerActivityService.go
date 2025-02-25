// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProviderActivityService is an autogenerated mock type for the providerActivityService type
type ProviderActivityService struct {
	mock.Mock
}

type ProviderActivityService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderActivityService) EXPECT() *ProviderActivityService_Expecter {
	return &ProviderActivityService_Expecter{mock: &_m.Mock}
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *ProviderActivityService) Find(_a0 context.Context, _a1 domain.ListProviderActivitiesFilter) ([]*domain.Activity, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Activity
	if rf, ok := ret.Get(0).(func(context.Context, domain.ListProviderActivitiesFilter) []*domain.Activity); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.ListProviderActivitiesFilter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderActivityService_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type ProviderActivityService_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 domain.ListProviderActivitiesFilter
func (_e *ProviderActivityService_Expecter) Find(_a0 interface{}, _a1 interface{}) *ProviderActivityService_Find_Call {
	return &ProviderActivityService_Find_Call{Call: _e.mock.On("Find", _a0, _a1)}
}

func (_c *ProviderActivityService_Find_Call) Run(run func(_a0 context.Context, _a1 domain.ListProviderActivitiesFilter)) *ProviderActivityService_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ListProviderActivitiesFilter))
	})
	return _c
}

func (_c *ProviderActivityService_Find_Call) Return(_a0 []*domain.Activity, _a1 error) *ProviderActivityService_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOne provides a mock function with given fields: _a0, _a1
func (_m *ProviderActivityService) GetOne(_a0 context.Context, _a1 string) (*domain.Activity, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Activity
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Activity); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderActivityService_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type ProviderActivityService_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 string
func (_e *ProviderActivityService_Expecter) GetOne(_a0 interface{}, _a1 interface{}) *ProviderActivityService_GetOne_Call {
	return &ProviderActivityService_GetOne_Call{Call: _e.mock.On("GetOne", _a0, _a1)}
}

func (_c *ProviderActivityService_GetOne_Call) Run(run func(_a0 context.Context, _a1 string)) *ProviderActivityService_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProviderActivityService_GetOne_Call) Return(_a0 *domain.Activity, _a1 error) *ProviderActivityService_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Import provides a mock function with given fields: _a0, _a1
func (_m *ProviderActivityService) Import(_a0 context.Context, _a1 domain.ImportActivitiesFilter) ([]*domain.Activity, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Activity
	if rf, ok := ret.Get(0).(func(context.Context, domain.ImportActivitiesFilter) []*domain.Activity); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.ImportActivitiesFilter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderActivityService_Import_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Import'
type ProviderActivityService_Import_Call struct {
	*mock.Call
}

// Import is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 domain.ImportActivitiesFilter
func (_e *ProviderActivityService_Expecter) Import(_a0 interface{}, _a1 interface{}) *ProviderActivityService_Import_Call {
	return &ProviderActivityService_Import_Call{Call: _e.mock.On("Import", _a0, _a1)}
}

func (_c *ProviderActivityService_Import_Call) Run(run func(_a0 context.Context, _a1 domain.ImportActivitiesFilter)) *ProviderActivityService_Import_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ImportActivitiesFilter))
	})
	return _c
}

func (_c *ProviderActivityService_Import_Call) Return(_a0 []*domain.Activity, _a1 error) *ProviderActivityService_Import_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

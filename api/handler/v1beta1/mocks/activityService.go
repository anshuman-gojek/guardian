// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// ActivityService is an autogenerated mock type for the activityService type
type ActivityService struct {
	mock.Mock
}

type ActivityService_Expecter struct {
	mock *mock.Mock
}

func (_m *ActivityService) EXPECT() *ActivityService_Expecter {
	return &ActivityService_Expecter{mock: &_m.Mock}
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *ActivityService) Find(_a0 context.Context, _a1 domain.ListProviderActivitiesFilter) ([]*domain.Activity, error) {
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

// ActivityService_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type ActivityService_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 domain.ListProviderActivitiesFilter
func (_e *ActivityService_Expecter) Find(_a0 interface{}, _a1 interface{}) *ActivityService_Find_Call {
	return &ActivityService_Find_Call{Call: _e.mock.On("Find", _a0, _a1)}
}

func (_c *ActivityService_Find_Call) Run(run func(_a0 context.Context, _a1 domain.ListProviderActivitiesFilter)) *ActivityService_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ListProviderActivitiesFilter))
	})
	return _c
}

func (_c *ActivityService_Find_Call) Return(_a0 []*domain.Activity, _a1 error) *ActivityService_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOne provides a mock function with given fields: _a0, _a1
func (_m *ActivityService) GetOne(_a0 context.Context, _a1 string) (*domain.Activity, error) {
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

// ActivityService_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type ActivityService_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 string
func (_e *ActivityService_Expecter) GetOne(_a0 interface{}, _a1 interface{}) *ActivityService_GetOne_Call {
	return &ActivityService_GetOne_Call{Call: _e.mock.On("GetOne", _a0, _a1)}
}

func (_c *ActivityService_GetOne_Call) Run(run func(_a0 context.Context, _a1 string)) *ActivityService_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ActivityService_GetOne_Call) Return(_a0 *domain.Activity, _a1 error) *ActivityService_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Import provides a mock function with given fields: _a0, _a1
func (_m *ActivityService) Import(_a0 context.Context, _a1 domain.ImportActivitiesFilter) ([]*domain.Activity, error) {
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

// ActivityService_Import_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Import'
type ActivityService_Import_Call struct {
	*mock.Call
}

// Import is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 domain.ImportActivitiesFilter
func (_e *ActivityService_Expecter) Import(_a0 interface{}, _a1 interface{}) *ActivityService_Import_Call {
	return &ActivityService_Import_Call{Call: _e.mock.On("Import", _a0, _a1)}
}

func (_c *ActivityService_Import_Call) Run(run func(_a0 context.Context, _a1 domain.ImportActivitiesFilter)) *ActivityService_Import_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ImportActivitiesFilter))
	})
	return _c
}

func (_c *ActivityService_Import_Call) Return(_a0 []*domain.Activity, _a1 error) *ActivityService_Import_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

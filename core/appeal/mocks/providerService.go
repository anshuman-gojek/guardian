// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProviderService is an autogenerated mock type for the providerService type
type ProviderService struct {
	mock.Mock
}

type ProviderService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderService) EXPECT() *ProviderService_Expecter {
	return &ProviderService_Expecter{mock: &_m.Mock}
}

// Find provides a mock function with given fields: _a0
func (_m *ProviderService) Find(_a0 context.Context) ([]*domain.Provider, error) {
	ret := _m.Called(_a0)

	var r0 []*domain.Provider
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Provider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderService_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type ProviderService_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *ProviderService_Expecter) Find(_a0 interface{}) *ProviderService_Find_Call {
	return &ProviderService_Find_Call{Call: _e.mock.On("Find", _a0)}
}

func (_c *ProviderService_Find_Call) Run(run func(_a0 context.Context)) *ProviderService_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ProviderService_Find_Call) Return(_a0 []*domain.Provider, _a1 error) *ProviderService_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPermissions provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *ProviderService) GetPermissions(_a0 context.Context, _a1 *domain.ProviderConfig, _a2 string, _a3 string) ([]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ProviderConfig, string, string) []interface{}); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ProviderConfig, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderService_GetPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissions'
type ProviderService_GetPermissions_Call struct {
	*mock.Call
}

// GetPermissions is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *domain.ProviderConfig
//  - _a2 string
//  - _a3 string
func (_e *ProviderService_Expecter) GetPermissions(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *ProviderService_GetPermissions_Call {
	return &ProviderService_GetPermissions_Call{Call: _e.mock.On("GetPermissions", _a0, _a1, _a2, _a3)}
}

func (_c *ProviderService_GetPermissions_Call) Run(run func(_a0 context.Context, _a1 *domain.ProviderConfig, _a2 string, _a3 string)) *ProviderService_GetPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ProviderConfig), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ProviderService_GetPermissions_Call) Return(_a0 []interface{}, _a1 error) *ProviderService_GetPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GrantAccess provides a mock function with given fields: _a0, _a1
func (_m *ProviderService) GrantAccess(_a0 context.Context, _a1 domain.Grant) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Grant) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProviderService_GrantAccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GrantAccess'
type ProviderService_GrantAccess_Call struct {
	*mock.Call
}

// GrantAccess is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 domain.Grant
func (_e *ProviderService_Expecter) GrantAccess(_a0 interface{}, _a1 interface{}) *ProviderService_GrantAccess_Call {
	return &ProviderService_GrantAccess_Call{Call: _e.mock.On("GrantAccess", _a0, _a1)}
}

func (_c *ProviderService_GrantAccess_Call) Run(run func(_a0 context.Context, _a1 domain.Grant)) *ProviderService_GrantAccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Grant))
	})
	return _c
}

func (_c *ProviderService_GrantAccess_Call) Return(_a0 error) *ProviderService_GrantAccess_Call {
	_c.Call.Return(_a0)
	return _c
}

// RevokeAccess provides a mock function with given fields: _a0, _a1
func (_m *ProviderService) RevokeAccess(_a0 context.Context, _a1 domain.Grant) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Grant) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProviderService_RevokeAccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeAccess'
type ProviderService_RevokeAccess_Call struct {
	*mock.Call
}

// RevokeAccess is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 domain.Grant
func (_e *ProviderService_Expecter) RevokeAccess(_a0 interface{}, _a1 interface{}) *ProviderService_RevokeAccess_Call {
	return &ProviderService_RevokeAccess_Call{Call: _e.mock.On("RevokeAccess", _a0, _a1)}
}

func (_c *ProviderService_RevokeAccess_Call) Run(run func(_a0 context.Context, _a1 domain.Grant)) *ProviderService_RevokeAccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Grant))
	})
	return _c
}

func (_c *ProviderService_RevokeAccess_Call) Return(_a0 error) *ProviderService_RevokeAccess_Call {
	_c.Call.Return(_a0)
	return _c
}

// ValidateAppeal provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *ProviderService) ValidateAppeal(_a0 context.Context, _a1 *domain.Appeal, _a2 *domain.Provider, _a3 *domain.Policy) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Appeal, *domain.Provider, *domain.Policy) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProviderService_ValidateAppeal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateAppeal'
type ProviderService_ValidateAppeal_Call struct {
	*mock.Call
}

// ValidateAppeal is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *domain.Appeal
//  - _a2 *domain.Provider
//  - _a3 *domain.Policy
func (_e *ProviderService_Expecter) ValidateAppeal(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *ProviderService_ValidateAppeal_Call {
	return &ProviderService_ValidateAppeal_Call{Call: _e.mock.On("ValidateAppeal", _a0, _a1, _a2, _a3)}
}

func (_c *ProviderService_ValidateAppeal_Call) Run(run func(_a0 context.Context, _a1 *domain.Appeal, _a2 *domain.Provider, _a3 *domain.Policy)) *ProviderService_ValidateAppeal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Appeal), args[2].(*domain.Provider), args[3].(*domain.Policy))
	})
	return _c
}

func (_c *ProviderService_ValidateAppeal_Call) Return(_a0 error) *ProviderService_ValidateAppeal_Call {
	_c.Call.Return(_a0)
	return _c
}

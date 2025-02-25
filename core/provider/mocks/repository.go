// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Repository) Create(_a0 context.Context, _a1 *domain.Provider) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Provider) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Repository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *domain.Provider
func (_e *Repository_Expecter) Create(_a0 interface{}, _a1 interface{}) *Repository_Create_Call {
	return &Repository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *Repository_Create_Call) Run(run func(_a0 context.Context, _a1 *domain.Provider)) *Repository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Provider))
	})
	return _c
}

func (_c *Repository_Create_Call) Return(_a0 error) *Repository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Repository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - id string
func (_e *Repository_Expecter) Delete(ctx interface{}, id interface{}) *Repository_Delete_Call {
	return &Repository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *Repository_Delete_Call) Run(run func(ctx context.Context, id string)) *Repository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_Delete_Call) Return(_a0 error) *Repository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Find provides a mock function with given fields: _a0
func (_m *Repository) Find(_a0 context.Context) ([]*domain.Provider, error) {
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

// Repository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type Repository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *Repository_Expecter) Find(_a0 interface{}) *Repository_Find_Call {
	return &Repository_Find_Call{Call: _e.mock.On("Find", _a0)}
}

func (_c *Repository_Find_Call) Run(run func(_a0 context.Context)) *Repository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Repository_Find_Call) Return(_a0 []*domain.Provider, _a1 error) *Repository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id string) (*domain.Provider, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Provider
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Provider); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type Repository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id string
func (_e *Repository_Expecter) GetByID(ctx interface{}, id interface{}) *Repository_GetByID_Call {
	return &Repository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *Repository_GetByID_Call) Run(run func(ctx context.Context, id string)) *Repository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_GetByID_Call) Return(_a0 *domain.Provider, _a1 error) *Repository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOne provides a mock function with given fields: ctx, pType, urn
func (_m *Repository) GetOne(ctx context.Context, pType string, urn string) (*domain.Provider, error) {
	ret := _m.Called(ctx, pType, urn)

	var r0 *domain.Provider
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.Provider); ok {
		r0 = rf(ctx, pType, urn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, pType, urn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type Repository_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//  - ctx context.Context
//  - pType string
//  - urn string
func (_e *Repository_Expecter) GetOne(ctx interface{}, pType interface{}, urn interface{}) *Repository_GetOne_Call {
	return &Repository_GetOne_Call{Call: _e.mock.On("GetOne", ctx, pType, urn)}
}

func (_c *Repository_GetOne_Call) Run(run func(ctx context.Context, pType string, urn string)) *Repository_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Repository_GetOne_Call) Return(_a0 *domain.Provider, _a1 error) *Repository_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTypes provides a mock function with given fields: _a0
func (_m *Repository) GetTypes(_a0 context.Context) ([]domain.ProviderType, error) {
	ret := _m.Called(_a0)

	var r0 []domain.ProviderType
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ProviderType); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProviderType)
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

// Repository_GetTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTypes'
type Repository_GetTypes_Call struct {
	*mock.Call
}

// GetTypes is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *Repository_Expecter) GetTypes(_a0 interface{}) *Repository_GetTypes_Call {
	return &Repository_GetTypes_Call{Call: _e.mock.On("GetTypes", _a0)}
}

func (_c *Repository_GetTypes_Call) Run(run func(_a0 context.Context)) *Repository_GetTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Repository_GetTypes_Call) Return(_a0 []domain.ProviderType, _a1 error) *Repository_GetTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Repository) Update(_a0 context.Context, _a1 *domain.Provider) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Provider) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Repository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *domain.Provider
func (_e *Repository_Expecter) Update(_a0 interface{}, _a1 interface{}) *Repository_Update_Call {
	return &Repository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *Repository_Update_Call) Run(run func(_a0 context.Context, _a1 *domain.Provider)) *Repository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Provider))
	})
	return _c
}

func (_c *Repository_Update_Call) Return(_a0 error) *Repository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

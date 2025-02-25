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

// BulkUpsert provides a mock function with given fields: _a0, _a1
func (_m *Repository) BulkUpsert(_a0 context.Context, _a1 []*domain.Appeal) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*domain.Appeal) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_BulkUpsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkUpsert'
type Repository_BulkUpsert_Call struct {
	*mock.Call
}

// BulkUpsert is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 []*domain.Appeal
func (_e *Repository_Expecter) BulkUpsert(_a0 interface{}, _a1 interface{}) *Repository_BulkUpsert_Call {
	return &Repository_BulkUpsert_Call{Call: _e.mock.On("BulkUpsert", _a0, _a1)}
}

func (_c *Repository_BulkUpsert_Call) Run(run func(_a0 context.Context, _a1 []*domain.Appeal)) *Repository_BulkUpsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*domain.Appeal))
	})
	return _c
}

func (_c *Repository_BulkUpsert_Call) Return(_a0 error) *Repository_BulkUpsert_Call {
	_c.Call.Return(_a0)
	return _c
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *Repository) Find(_a0 context.Context, _a1 *domain.ListAppealsFilter) ([]*domain.Appeal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Appeal
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ListAppealsFilter) []*domain.Appeal); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Appeal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ListAppealsFilter) error); ok {
		r1 = rf(_a0, _a1)
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
//  - _a1 *domain.ListAppealsFilter
func (_e *Repository_Expecter) Find(_a0 interface{}, _a1 interface{}) *Repository_Find_Call {
	return &Repository_Find_Call{Call: _e.mock.On("Find", _a0, _a1)}
}

func (_c *Repository_Find_Call) Run(run func(_a0 context.Context, _a1 *domain.ListAppealsFilter)) *Repository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ListAppealsFilter))
	})
	return _c
}

func (_c *Repository_Find_Call) Return(_a0 []*domain.Appeal, _a1 error) *Repository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id string) (*domain.Appeal, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Appeal
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Appeal); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Appeal)
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

func (_c *Repository_GetByID_Call) Return(_a0 *domain.Appeal, _a1 error) *Repository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Repository) Update(_a0 context.Context, _a1 *domain.Appeal) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Appeal) error); ok {
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
//  - _a1 *domain.Appeal
func (_e *Repository_Expecter) Update(_a0 interface{}, _a1 interface{}) *Repository_Update_Call {
	return &Repository_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *Repository_Update_Call) Run(run func(_a0 context.Context, _a1 *domain.Appeal)) *Repository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Appeal))
	})
	return _c
}

func (_c *Repository_Update_Call) Return(_a0 error) *Repository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

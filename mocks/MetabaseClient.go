// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	metabase "github.com/odpf/guardian/plugins/providers/metabase"
	mock "github.com/stretchr/testify/mock"
)

// MetabaseClient is an autogenerated mock type for the MetabaseClient type
type MetabaseClient struct {
	mock.Mock
}

// GetCollections provides a mock function with given fields:
func (_m *MetabaseClient) GetCollections() ([]*metabase.Collection, error) {
	ret := _m.Called()

	var r0 []*metabase.Collection
	if rf, ok := ret.Get(0).(func() []*metabase.Collection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*metabase.Collection)
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

// GetDatabases provides a mock function with given fields:
func (_m *MetabaseClient) GetDatabases() ([]*metabase.Database, error) {
	ret := _m.Called()

	var r0 []*metabase.Database
	if rf, ok := ret.Get(0).(func() []*metabase.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*metabase.Database)
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

// GetGroups provides a mock function with given fields:
func (_m *MetabaseClient) GetGroups() ([]*metabase.Group, metabase.ResourceGroupDetails, metabase.ResourceGroupDetails, error) {
	ret := _m.Called()

	var r0 []*metabase.Group
	if rf, ok := ret.Get(0).(func() []*metabase.Group); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*metabase.Group)
		}
	}

	var r1 metabase.ResourceGroupDetails
	if rf, ok := ret.Get(1).(func() metabase.ResourceGroupDetails); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(metabase.ResourceGroupDetails)
		}
	}

	var r2 metabase.ResourceGroupDetails
	if rf, ok := ret.Get(2).(func() metabase.ResourceGroupDetails); ok {
		r2 = rf()
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(metabase.ResourceGroupDetails)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func() error); ok {
		r3 = rf()
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GrantCollectionAccess provides a mock function with given fields: resource, user, role
func (_m *MetabaseClient) GrantCollectionAccess(resource *metabase.Collection, user string, role string) error {
	ret := _m.Called(resource, user, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metabase.Collection, string, string) error); ok {
		r0 = rf(resource, user, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantDatabaseAccess provides a mock function with given fields: resource, user, role, groups
func (_m *MetabaseClient) GrantDatabaseAccess(resource *metabase.Database, user string, role string, groups map[string]*metabase.Group) error {
	ret := _m.Called(resource, user, role, groups)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metabase.Database, string, string, map[string]*metabase.Group) error); ok {
		r0 = rf(resource, user, role, groups)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantGroupAccess provides a mock function with given fields: groupID, email
func (_m *MetabaseClient) GrantGroupAccess(groupID int, email string) error {
	ret := _m.Called(groupID, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(groupID, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantTableAccess provides a mock function with given fields: resource, user, role, groups
func (_m *MetabaseClient) GrantTableAccess(resource *metabase.Table, user string, role string, groups map[string]*metabase.Group) error {
	ret := _m.Called(resource, user, role, groups)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metabase.Table, string, string, map[string]*metabase.Group) error); ok {
		r0 = rf(resource, user, role, groups)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeCollectionAccess provides a mock function with given fields: resource, user, role
func (_m *MetabaseClient) RevokeCollectionAccess(resource *metabase.Collection, user string, role string) error {
	ret := _m.Called(resource, user, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metabase.Collection, string, string) error); ok {
		r0 = rf(resource, user, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeDatabaseAccess provides a mock function with given fields: resource, user, role
func (_m *MetabaseClient) RevokeDatabaseAccess(resource *metabase.Database, user string, role string) error {
	ret := _m.Called(resource, user, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metabase.Database, string, string) error); ok {
		r0 = rf(resource, user, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeGroupAccess provides a mock function with given fields: groupID, email
func (_m *MetabaseClient) RevokeGroupAccess(groupID int, email string) error {
	ret := _m.Called(groupID, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(groupID, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeTableAccess provides a mock function with given fields: resource, user, role
func (_m *MetabaseClient) RevokeTableAccess(resource *metabase.Table, user string, role string) error {
	ret := _m.Called(resource, user, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metabase.Table, string, string) error); ok {
		r0 = rf(resource, user, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

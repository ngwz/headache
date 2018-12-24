// Code generated by mockery v1.0.0. DO NOT EDIT.

package vcs_mocks

import helper "github.com/fbiville/headache/helper"
import mock "github.com/stretchr/testify/mock"
import vcs "github.com/fbiville/headache/vcs"

// VersioningClient is an autogenerated mock type for the VersioningClient type
type VersioningClient struct {
	mock.Mock
}

// AddMetadata provides a mock function with given fields: changes, clock
func (_m *VersioningClient) AddMetadata(changes []vcs.FileChange, clock helper.Clock) ([]vcs.FileChange, error) {
	ret := _m.Called(changes, clock)

	var r0 []vcs.FileChange
	if rf, ok := ret.Get(0).(func([]vcs.FileChange, helper.Clock) []vcs.FileChange); ok {
		r0 = rf(changes, clock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vcs.FileChange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]vcs.FileChange, helper.Clock) error); ok {
		r1 = rf(changes, clock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChanges provides a mock function with given fields: revision
func (_m *VersioningClient) GetChanges(revision string) ([]vcs.FileChange, error) {
	ret := _m.Called(revision)

	var r0 []vcs.FileChange
	if rf, ok := ret.Get(0).(func(string) []vcs.FileChange); ok {
		r0 = rf(revision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vcs.FileChange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClient provides a mock function with given fields:
func (_m *VersioningClient) GetClient() vcs.Vcs {
	ret := _m.Called()

	var r0 vcs.Vcs
	if rf, ok := ret.Get(0).(func() vcs.Vcs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vcs.Vcs)
		}
	}

	return r0
}

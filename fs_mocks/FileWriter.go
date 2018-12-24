// Code generated by mockery v1.0.0. DO NOT EDIT.

package fs_mocks

import fs "github.com/fbiville/headache/fs"
import mock "github.com/stretchr/testify/mock"
import os "os"

// FileWriter is an autogenerated mock type for the FileWriter type
type FileWriter struct {
	mock.Mock
}

// Open provides a mock function with given fields: path, mask, permissions
func (_m *FileWriter) Open(path string, mask int, permissions os.FileMode) (fs.File, error) {
	ret := _m.Called(path, mask, permissions)

	var r0 fs.File
	if rf, ok := ret.Get(0).(func(string, int, os.FileMode) fs.File); ok {
		r0 = rf(path, mask, permissions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, os.FileMode) error); ok {
		r1 = rf(path, mask, permissions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: path, contents, permissions
func (_m *FileWriter) Write(path string, contents string, permissions os.FileMode) error {
	ret := _m.Called(path, contents, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, os.FileMode) error); ok {
		r0 = rf(path, contents, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	state "github.com/franono/tendermint/state"

	types "github.com/franono/tendermint/types"
)

// StateProvider is an autogenerated mock type for the StateProvider type
type StateProvider struct {
	mock.Mock
}

// AppHash provides a mock function with given fields: height
func (_m *StateProvider) AppHash(height uint64) ([]byte, error) {
	ret := _m.Called(height)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(uint64) []byte); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: height
func (_m *StateProvider) Commit(height uint64) (*types.Commit, error) {
	ret := _m.Called(height)

	var r0 *types.Commit
	if rf, ok := ret.Get(0).(func(uint64) *types.Commit); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Commit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// State provides a mock function with given fields: height
func (_m *StateProvider) State(height uint64) (state.State, error) {
	ret := _m.Called(height)

	var r0 state.State
	if rf, ok := ret.Get(0).(func(uint64) state.State); ok {
		r0 = rf(height)
	} else {
		r0 = ret.Get(0).(state.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

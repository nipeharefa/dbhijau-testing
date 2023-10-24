// Code generated by MockGen. DO NOT EDIT.
// Source: demo.nipeharefa.dev/demo (interfaces: Finder)
//
// Generated by this command:
//
//	mockgen --typed --package=mocks --destination=mocks/finder.go demo.nipeharefa.dev/demo Finder
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFinder is a mock of Finder interface.
type MockFinder struct {
	ctrl     *gomock.Controller
	recorder *MockFinderMockRecorder
}

// MockFinderMockRecorder is the mock recorder for MockFinder.
type MockFinderMockRecorder struct {
	mock *MockFinder
}

// NewMockFinder creates a new mock instance.
func NewMockFinder(ctrl *gomock.Controller) *MockFinder {
	mock := &MockFinder{ctrl: ctrl}
	mock.recorder = &MockFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinder) EXPECT() *MockFinderMockRecorder {
	return m.recorder
}

// Say mocks base method.
func (m *MockFinder) Say() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Say")
	ret0, _ := ret[0].(string)
	return ret0
}

// Say indicates an expected call of Say.
func (mr *MockFinderMockRecorder) Say() *FinderSayCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Say", reflect.TypeOf((*MockFinder)(nil).Say))
	return &FinderSayCall{Call: call}
}

// FinderSayCall wrap *gomock.Call
type FinderSayCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *FinderSayCall) Return(arg0 string) *FinderSayCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *FinderSayCall) Do(f func() string) *FinderSayCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *FinderSayCall) DoAndReturn(f func() string) *FinderSayCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

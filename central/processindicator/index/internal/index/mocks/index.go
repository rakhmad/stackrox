// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/processindicator/index/internal/index (interfaces: Indexer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	reflect "reflect"
)

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddProcessIndicator mocks base method
func (m *MockIndexer) AddProcessIndicator(arg0 *storage.ProcessIndicator) error {
	ret := m.ctrl.Call(m, "AddProcessIndicator", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessIndicator indicates an expected call of AddProcessIndicator
func (mr *MockIndexerMockRecorder) AddProcessIndicator(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessIndicator", reflect.TypeOf((*MockIndexer)(nil).AddProcessIndicator), arg0)
}

// AddProcessIndicators mocks base method
func (m *MockIndexer) AddProcessIndicators(arg0 []*storage.ProcessIndicator) error {
	ret := m.ctrl.Call(m, "AddProcessIndicators", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessIndicators indicates an expected call of AddProcessIndicators
func (mr *MockIndexerMockRecorder) AddProcessIndicators(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessIndicators", reflect.TypeOf((*MockIndexer)(nil).AddProcessIndicators), arg0)
}

// DeleteProcessIndicator mocks base method
func (m *MockIndexer) DeleteProcessIndicator(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteProcessIndicator", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProcessIndicator indicates an expected call of DeleteProcessIndicator
func (mr *MockIndexerMockRecorder) DeleteProcessIndicator(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProcessIndicator", reflect.TypeOf((*MockIndexer)(nil).DeleteProcessIndicator), arg0)
}

// Search mocks base method
func (m *MockIndexer) Search(arg0 *v1.Query) ([]search.Result, error) {
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIndexerMockRecorder) Search(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), arg0)
}

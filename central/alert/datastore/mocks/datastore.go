// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/alert/datastore (interfaces: DataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddAlert mocks base method
func (m *MockDataStore) AddAlert(arg0 *storage.Alert) error {
	ret := m.ctrl.Call(m, "AddAlert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAlert indicates an expected call of AddAlert
func (mr *MockDataStoreMockRecorder) AddAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlert", reflect.TypeOf((*MockDataStore)(nil).AddAlert), arg0)
}

// CountAlerts mocks base method
func (m *MockDataStore) CountAlerts() (int, error) {
	ret := m.ctrl.Call(m, "CountAlerts")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAlerts indicates an expected call of CountAlerts
func (mr *MockDataStoreMockRecorder) CountAlerts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAlerts", reflect.TypeOf((*MockDataStore)(nil).CountAlerts))
}

// GetAlert mocks base method
func (m *MockDataStore) GetAlert(arg0 string) (*storage.Alert, bool, error) {
	ret := m.ctrl.Call(m, "GetAlert", arg0)
	ret0, _ := ret[0].(*storage.Alert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlert indicates an expected call of GetAlert
func (mr *MockDataStoreMockRecorder) GetAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlert", reflect.TypeOf((*MockDataStore)(nil).GetAlert), arg0)
}

// GetAlertStore mocks base method
func (m *MockDataStore) GetAlertStore() ([]*storage.ListAlert, error) {
	ret := m.ctrl.Call(m, "GetAlertStore")
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertStore indicates an expected call of GetAlertStore
func (mr *MockDataStoreMockRecorder) GetAlertStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertStore", reflect.TypeOf((*MockDataStore)(nil).GetAlertStore))
}

// ListAlerts mocks base method
func (m *MockDataStore) ListAlerts(arg0 *v1.ListAlertsRequest) ([]*storage.ListAlert, error) {
	ret := m.ctrl.Call(m, "ListAlerts", arg0)
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlerts indicates an expected call of ListAlerts
func (mr *MockDataStoreMockRecorder) ListAlerts(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlerts", reflect.TypeOf((*MockDataStore)(nil).ListAlerts), arg0)
}

// MarkAlertStale mocks base method
func (m *MockDataStore) MarkAlertStale(arg0 string) error {
	ret := m.ctrl.Call(m, "MarkAlertStale", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAlertStale indicates an expected call of MarkAlertStale
func (mr *MockDataStoreMockRecorder) MarkAlertStale(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAlertStale", reflect.TypeOf((*MockDataStore)(nil).MarkAlertStale), arg0)
}

// SearchAlerts mocks base method
func (m *MockDataStore) SearchAlerts(arg0 *v1.Query) ([]*v1.SearchResult, error) {
	ret := m.ctrl.Call(m, "SearchAlerts", arg0)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAlerts indicates an expected call of SearchAlerts
func (mr *MockDataStoreMockRecorder) SearchAlerts(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchAlerts), arg0)
}

// SearchListAlerts mocks base method
func (m *MockDataStore) SearchListAlerts(arg0 *v1.Query) ([]*storage.ListAlert, error) {
	ret := m.ctrl.Call(m, "SearchListAlerts", arg0)
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListAlerts indicates an expected call of SearchListAlerts
func (mr *MockDataStoreMockRecorder) SearchListAlerts(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchListAlerts), arg0)
}

// SearchRawAlerts mocks base method
func (m *MockDataStore) SearchRawAlerts(arg0 *v1.Query) ([]*storage.Alert, error) {
	ret := m.ctrl.Call(m, "SearchRawAlerts", arg0)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawAlerts indicates an expected call of SearchRawAlerts
func (mr *MockDataStoreMockRecorder) SearchRawAlerts(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchRawAlerts), arg0)
}

// UpdateAlert mocks base method
func (m *MockDataStore) UpdateAlert(arg0 *storage.Alert) error {
	ret := m.ctrl.Call(m, "UpdateAlert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlert indicates an expected call of UpdateAlert
func (mr *MockDataStoreMockRecorder) UpdateAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlert", reflect.TypeOf((*MockDataStore)(nil).UpdateAlert), arg0)
}

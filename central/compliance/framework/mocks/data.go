// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/compliance/framework (interfaces: ComplianceDataRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	compliance "github.com/stackrox/rox/generated/internalapi/compliance"
	storage "github.com/stackrox/rox/generated/storage"
	set "github.com/stackrox/rox/pkg/set"
	reflect "reflect"
)

// MockComplianceDataRepository is a mock of ComplianceDataRepository interface
type MockComplianceDataRepository struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceDataRepositoryMockRecorder
}

// MockComplianceDataRepositoryMockRecorder is the mock recorder for MockComplianceDataRepository
type MockComplianceDataRepositoryMockRecorder struct {
	mock *MockComplianceDataRepository
}

// NewMockComplianceDataRepository creates a new mock instance
func NewMockComplianceDataRepository(ctrl *gomock.Controller) *MockComplianceDataRepository {
	mock := &MockComplianceDataRepository{ctrl: ctrl}
	mock.recorder = &MockComplianceDataRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockComplianceDataRepository) EXPECT() *MockComplianceDataRepositoryMockRecorder {
	return m.recorder
}

// Alerts mocks base method
func (m *MockComplianceDataRepository) Alerts() []*storage.ListAlert {
	ret := m.ctrl.Call(m, "Alerts")
	ret0, _ := ret[0].([]*storage.ListAlert)
	return ret0
}

// Alerts indicates an expected call of Alerts
func (mr *MockComplianceDataRepositoryMockRecorder) Alerts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alerts", reflect.TypeOf((*MockComplianceDataRepository)(nil).Alerts))
}

// Cluster mocks base method
func (m *MockComplianceDataRepository) Cluster() *storage.Cluster {
	ret := m.ctrl.Call(m, "Cluster")
	ret0, _ := ret[0].(*storage.Cluster)
	return ret0
}

// Cluster indicates an expected call of Cluster
func (mr *MockComplianceDataRepositoryMockRecorder) Cluster() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockComplianceDataRepository)(nil).Cluster))
}

// Deployments mocks base method
func (m *MockComplianceDataRepository) Deployments() map[string]*storage.Deployment {
	ret := m.ctrl.Call(m, "Deployments")
	ret0, _ := ret[0].(map[string]*storage.Deployment)
	return ret0
}

// Deployments indicates an expected call of Deployments
func (mr *MockComplianceDataRepositoryMockRecorder) Deployments() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockComplianceDataRepository)(nil).Deployments))
}

// HostFiles mocks base method
func (m *MockComplianceDataRepository) HostFiles(arg0 *storage.Node) map[string]*compliance.File {
	ret := m.ctrl.Call(m, "HostFiles", arg0)
	ret0, _ := ret[0].(map[string]*compliance.File)
	return ret0
}

// HostFiles indicates an expected call of HostFiles
func (mr *MockComplianceDataRepositoryMockRecorder) HostFiles(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostFiles", reflect.TypeOf((*MockComplianceDataRepository)(nil).HostFiles), arg0)
}

// HostProcesses mocks base method
func (m *MockComplianceDataRepository) HostProcesses(arg0 *storage.Node) []*compliance.CommandLine {
	ret := m.ctrl.Call(m, "HostProcesses", arg0)
	ret0, _ := ret[0].([]*compliance.CommandLine)
	return ret0
}

// HostProcesses indicates an expected call of HostProcesses
func (mr *MockComplianceDataRepositoryMockRecorder) HostProcesses(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostProcesses", reflect.TypeOf((*MockComplianceDataRepository)(nil).HostProcesses), arg0)
}

// ImageIntegrations mocks base method
func (m *MockComplianceDataRepository) ImageIntegrations() []*storage.ImageIntegration {
	ret := m.ctrl.Call(m, "ImageIntegrations")
	ret0, _ := ret[0].([]*storage.ImageIntegration)
	return ret0
}

// ImageIntegrations indicates an expected call of ImageIntegrations
func (mr *MockComplianceDataRepositoryMockRecorder) ImageIntegrations() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageIntegrations", reflect.TypeOf((*MockComplianceDataRepository)(nil).ImageIntegrations))
}

// NetworkFlows mocks base method
func (m *MockComplianceDataRepository) NetworkFlows() []*storage.NetworkFlow {
	ret := m.ctrl.Call(m, "NetworkFlows")
	ret0, _ := ret[0].([]*storage.NetworkFlow)
	return ret0
}

// NetworkFlows indicates an expected call of NetworkFlows
func (mr *MockComplianceDataRepositoryMockRecorder) NetworkFlows() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkFlows", reflect.TypeOf((*MockComplianceDataRepository)(nil).NetworkFlows))
}

// NetworkGraph mocks base method
func (m *MockComplianceDataRepository) NetworkGraph() *v1.NetworkGraph {
	ret := m.ctrl.Call(m, "NetworkGraph")
	ret0, _ := ret[0].(*v1.NetworkGraph)
	return ret0
}

// NetworkGraph indicates an expected call of NetworkGraph
func (mr *MockComplianceDataRepositoryMockRecorder) NetworkGraph() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkGraph", reflect.TypeOf((*MockComplianceDataRepository)(nil).NetworkGraph))
}

// NetworkPolicies mocks base method
func (m *MockComplianceDataRepository) NetworkPolicies() map[string]*storage.NetworkPolicy {
	ret := m.ctrl.Call(m, "NetworkPolicies")
	ret0, _ := ret[0].(map[string]*storage.NetworkPolicy)
	return ret0
}

// NetworkPolicies indicates an expected call of NetworkPolicies
func (mr *MockComplianceDataRepositoryMockRecorder) NetworkPolicies() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkPolicies", reflect.TypeOf((*MockComplianceDataRepository)(nil).NetworkPolicies))
}

// Nodes mocks base method
func (m *MockComplianceDataRepository) Nodes() map[string]*storage.Node {
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(map[string]*storage.Node)
	return ret0
}

// Nodes indicates an expected call of Nodes
func (mr *MockComplianceDataRepositoryMockRecorder) Nodes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockComplianceDataRepository)(nil).Nodes))
}

// Policies mocks base method
func (m *MockComplianceDataRepository) Policies() map[string]*storage.Policy {
	ret := m.ctrl.Call(m, "Policies")
	ret0, _ := ret[0].(map[string]*storage.Policy)
	return ret0
}

// Policies indicates an expected call of Policies
func (mr *MockComplianceDataRepositoryMockRecorder) Policies() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policies", reflect.TypeOf((*MockComplianceDataRepository)(nil).Policies))
}

// PolicyCategories mocks base method
func (m *MockComplianceDataRepository) PolicyCategories() map[string]set.StringSet {
	ret := m.ctrl.Call(m, "PolicyCategories")
	ret0, _ := ret[0].(map[string]set.StringSet)
	return ret0
}

// PolicyCategories indicates an expected call of PolicyCategories
func (mr *MockComplianceDataRepositoryMockRecorder) PolicyCategories() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyCategories", reflect.TypeOf((*MockComplianceDataRepository)(nil).PolicyCategories))
}

// ProcessIndicators mocks base method
func (m *MockComplianceDataRepository) ProcessIndicators() []*storage.ProcessIndicator {
	ret := m.ctrl.Call(m, "ProcessIndicators")
	ret0, _ := ret[0].([]*storage.ProcessIndicator)
	return ret0
}

// ProcessIndicators indicates an expected call of ProcessIndicators
func (mr *MockComplianceDataRepositoryMockRecorder) ProcessIndicators() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessIndicators", reflect.TypeOf((*MockComplianceDataRepository)(nil).ProcessIndicators))
}

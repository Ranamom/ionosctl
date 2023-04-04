// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/v6/services/dbaas-postgres/resources"
)

// MockClustersService is a mock of ClustersService interface.
type MockClustersService struct {
	ctrl     *gomock.Controller
	recorder *MockClustersServiceMockRecorder
}

// MockClustersServiceMockRecorder is the mock recorder for MockClustersService.
type MockClustersServiceMockRecorder struct {
	mock *MockClustersService
}

// NewMockClustersService creates a new mock instance.
func NewMockClustersService(ctrl *gomock.Controller) *MockClustersService {
	mock := &MockClustersService{ctrl: ctrl}
	mock.recorder = &MockClustersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClustersService) EXPECT() *MockClustersServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClustersService) Create(input resources.CreateClusterRequest) (*resources.ClusterResponse, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*resources.ClusterResponse)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockClustersServiceMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClustersService)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockClustersService) Delete(clusterId string) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", clusterId)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockClustersServiceMockRecorder) Delete(clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClustersService)(nil).Delete), clusterId)
}

// Get mocks base method.
func (m *MockClustersService) Get(clusterId string) (*resources.ClusterResponse, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", clusterId)
	ret0, _ := ret[0].(*resources.ClusterResponse)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockClustersServiceMockRecorder) Get(clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClustersService)(nil).Get), clusterId)
}

// List mocks base method.
func (m *MockClustersService) List(filterName string) (resources.ClusterList, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filterName)
	ret0, _ := ret[0].(resources.ClusterList)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockClustersServiceMockRecorder) List(filterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClustersService)(nil).List), filterName)
}

// Update mocks base method.
func (m *MockClustersService) Update(clusterId string, input resources.PatchClusterRequest) (*resources.ClusterResponse, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", clusterId, input)
	ret0, _ := ret[0].(*resources.ClusterResponse)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockClustersServiceMockRecorder) Update(clusterId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClustersService)(nil).Update), clusterId, input)
}

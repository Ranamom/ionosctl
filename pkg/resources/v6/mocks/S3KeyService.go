// Code generated by MockGen. DO NOT EDIT.
// Source: s3key.go

// Package mock_v6 is a generated GoMock package.
package mock_v6

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v6 "github.com/ionos-cloud/ionosctl/pkg/resources/v6"
)

// MockS3KeysService is a mock of S3KeysService interface.
type MockS3KeysService struct {
	ctrl     *gomock.Controller
	recorder *MockS3KeysServiceMockRecorder
}

// MockS3KeysServiceMockRecorder is the mock recorder for MockS3KeysService.
type MockS3KeysServiceMockRecorder struct {
	mock *MockS3KeysService
}

// NewMockS3KeysService creates a new mock instance.
func NewMockS3KeysService(ctrl *gomock.Controller) *MockS3KeysService {
	mock := &MockS3KeysService{ctrl: ctrl}
	mock.recorder = &MockS3KeysServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3KeysService) EXPECT() *MockS3KeysServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockS3KeysService) Create(userId string) (*v6.S3Key, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId)
	ret0, _ := ret[0].(*v6.S3Key)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockS3KeysServiceMockRecorder) Create(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockS3KeysService)(nil).Create), userId)
}

// Delete mocks base method.
func (m *MockS3KeysService) Delete(userId, keyId string) (*v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, keyId)
	ret0, _ := ret[0].(*v6.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockS3KeysServiceMockRecorder) Delete(userId, keyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockS3KeysService)(nil).Delete), userId, keyId)
}

// Get mocks base method.
func (m *MockS3KeysService) Get(userId, keyId string) (*v6.S3Key, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userId, keyId)
	ret0, _ := ret[0].(*v6.S3Key)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockS3KeysServiceMockRecorder) Get(userId, keyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockS3KeysService)(nil).Get), userId, keyId)
}

// List mocks base method.
func (m *MockS3KeysService) List(userId string) (v6.S3Keys, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userId)
	ret0, _ := ret[0].(v6.S3Keys)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockS3KeysServiceMockRecorder) List(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockS3KeysService)(nil).List), userId)
}

// Update mocks base method.
func (m *MockS3KeysService) Update(userId, keyId string, key v6.S3Key) (*v6.S3Key, *v6.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, keyId, key)
	ret0, _ := ret[0].(*v6.S3Key)
	ret1, _ := ret[1].(*v6.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockS3KeysServiceMockRecorder) Update(userId, keyId, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockS3KeysService)(nil).Update), userId, keyId, key)
}
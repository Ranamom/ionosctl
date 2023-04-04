// Code generated by MockGen. DO NOT EDIT.
// Source: template.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/v6/services/cloudapi-v6/resources"
)

// MockTemplatesService is a mock of TemplatesService interface.
type MockTemplatesService struct {
	ctrl     *gomock.Controller
	recorder *MockTemplatesServiceMockRecorder
}

// MockTemplatesServiceMockRecorder is the mock recorder for MockTemplatesService.
type MockTemplatesServiceMockRecorder struct {
	mock *MockTemplatesService
}

// NewMockTemplatesService creates a new mock instance.
func NewMockTemplatesService(ctrl *gomock.Controller) *MockTemplatesService {
	mock := &MockTemplatesService{ctrl: ctrl}
	mock.recorder = &MockTemplatesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemplatesService) EXPECT() *MockTemplatesServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTemplatesService) Get(templateId string, params resources.QueryParams) (*resources.Template, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", templateId, params)
	ret0, _ := ret[0].(*resources.Template)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockTemplatesServiceMockRecorder) Get(templateId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTemplatesService)(nil).Get), templateId, params)
}

// List mocks base method.
func (m *MockTemplatesService) List(params resources.ListQueryParams) (resources.Templates, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", params)
	ret0, _ := ret[0].(resources.Templates)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockTemplatesServiceMockRecorder) List(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTemplatesService)(nil).List), params)
}

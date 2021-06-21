// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: BackupApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockBackupApi is a mock of BackupApi interface.
type MockBackupApi struct {
	ctrl     *gomock.Controller
	recorder *MockBackupApiMockRecorder
}

// MockBackupApiMockRecorder is the mock recorder for MockBackupApi.
type MockBackupApiMockRecorder struct {
	mock *MockBackupApi
}

// NewMockBackupApi creates a new mock instance.
func NewMockBackupApi(ctrl *gomock.Controller) *MockBackupApi {
	mock := &MockBackupApi{ctrl: ctrl}
	mock.recorder = &MockBackupApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupApi) EXPECT() *MockBackupApiMockRecorder {
	return m.recorder
}

// GetBackupMetadata mocks base method.
func (m *MockBackupApi) GetBackupMetadata(arg0 context.Context) api.ApiGetBackupMetadataRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackupMetadata", arg0)
	ret0, _ := ret[0].(api.ApiGetBackupMetadataRequest)
	return ret0
}

// GetBackupMetadata indicates an expected call of GetBackupMetadata.
func (mr *MockBackupApiMockRecorder) GetBackupMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupMetadata", reflect.TypeOf((*MockBackupApi)(nil).GetBackupMetadata), arg0)
}

// GetBackupMetadataExecute mocks base method.
func (m *MockBackupApi) GetBackupMetadataExecute(arg0 api.ApiGetBackupMetadataRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackupMetadataExecute", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupMetadataExecute indicates an expected call of GetBackupMetadataExecute.
func (mr *MockBackupApiMockRecorder) GetBackupMetadataExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupMetadataExecute", reflect.TypeOf((*MockBackupApi)(nil).GetBackupMetadataExecute), arg0)
}

// GetBackupShardId mocks base method.
func (m *MockBackupApi) GetBackupShardId(arg0 context.Context, arg1 int64) api.ApiGetBackupShardIdRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackupShardId", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetBackupShardIdRequest)
	return ret0
}

// GetBackupShardId indicates an expected call of GetBackupShardId.
func (mr *MockBackupApiMockRecorder) GetBackupShardId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupShardId", reflect.TypeOf((*MockBackupApi)(nil).GetBackupShardId), arg0, arg1)
}

// GetBackupShardIdExecute mocks base method.
func (m *MockBackupApi) GetBackupShardIdExecute(arg0 api.ApiGetBackupShardIdRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackupShardIdExecute", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupShardIdExecute indicates an expected call of GetBackupShardIdExecute.
func (mr *MockBackupApiMockRecorder) GetBackupShardIdExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupShardIdExecute", reflect.TypeOf((*MockBackupApi)(nil).GetBackupShardIdExecute), arg0)
}
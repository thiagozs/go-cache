// Code generated by MockGen. DO NOT EDIT.
// Source: buntdb/repository.go

// Package buntdb is a generated GoMock package.
package buntdb

import (
	reflect "reflect"

	kind "github.com/thiagozs/go-cache/kind"
	gomock "go.uber.org/mock/gomock"
)

// MockBuntDBLayerRepo is a mock of BuntDBLayerRepo interface.
type MockBuntDBLayerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBuntDBLayerRepoMockRecorder
}

// MockBuntDBLayerRepoMockRecorder is the mock recorder for MockBuntDBLayerRepo.
type MockBuntDBLayerRepoMockRecorder struct {
	mock *MockBuntDBLayerRepo
}

// NewMockBuntDBLayerRepo creates a new mock instance.
func NewMockBuntDBLayerRepo(ctrl *gomock.Controller) *MockBuntDBLayerRepo {
	mock := &MockBuntDBLayerRepo{ctrl: ctrl}
	mock.recorder = &MockBuntDBLayerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuntDBLayerRepo) EXPECT() *MockBuntDBLayerRepoMockRecorder {
	return m.recorder
}

// DeleteKey mocks base method.
func (m *MockBuntDBLayerRepo) DeleteKey(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockBuntDBLayerRepoMockRecorder) DeleteKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).DeleteKey), key)
}

// GetDriver mocks base method.
func (m *MockBuntDBLayerRepo) GetDriver() kind.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriver")
	ret0, _ := ret[0].(kind.Driver)
	return ret0
}

// GetDriver indicates an expected call of GetDriver.
func (mr *MockBuntDBLayerRepoMockRecorder) GetDriver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriver", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).GetDriver))
}

// GetVal mocks base method.
func (m *MockBuntDBLayerRepo) GetVal(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVal", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVal indicates an expected call of GetVal.
func (mr *MockBuntDBLayerRepoMockRecorder) GetVal(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVal", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).GetVal), key)
}

// WriteKeyVal mocks base method.
func (m *MockBuntDBLayerRepo) WriteKeyVal(key, val string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyVal", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyVal indicates an expected call of WriteKeyVal.
func (mr *MockBuntDBLayerRepoMockRecorder) WriteKeyVal(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyVal", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).WriteKeyVal), key, val)
}

// WriteKeyValAsJSON mocks base method.
func (m *MockBuntDBLayerRepo) WriteKeyValAsJSON(key string, val any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyValAsJSON", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyValAsJSON indicates an expected call of WriteKeyValAsJSON.
func (mr *MockBuntDBLayerRepoMockRecorder) WriteKeyValAsJSON(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyValAsJSON", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).WriteKeyValAsJSON), key, val)
}

// WriteKeyValAsJSONTTL mocks base method.
func (m *MockBuntDBLayerRepo) WriteKeyValAsJSONTTL(key string, val any, insec int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyValAsJSONTTL", key, val, insec)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyValAsJSONTTL indicates an expected call of WriteKeyValAsJSONTTL.
func (mr *MockBuntDBLayerRepoMockRecorder) WriteKeyValAsJSONTTL(key, val, insec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyValAsJSONTTL", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).WriteKeyValAsJSONTTL), key, val, insec)
}

// WriteKeyValTTL mocks base method.
func (m *MockBuntDBLayerRepo) WriteKeyValTTL(key, val string, insec int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyValTTL", key, val, insec)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyValTTL indicates an expected call of WriteKeyValTTL.
func (mr *MockBuntDBLayerRepoMockRecorder) WriteKeyValTTL(key, val, insec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyValTTL", reflect.TypeOf((*MockBuntDBLayerRepo)(nil).WriteKeyValTTL), key, val, insec)
}
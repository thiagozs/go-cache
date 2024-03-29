// Code generated by MockGen. DO NOT EDIT.
// Source: redis/repository.go

// Package redis is a generated GoMock package.
package redis

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kind "github.com/thiagozs/go-cache/kind"
)

// MockRedisLayerRepo is a mock of RedisLayerRepo interface.
type MockRedisLayerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRedisLayerRepoMockRecorder
}

// MockRedisLayerRepoMockRecorder is the mock recorder for MockRedisLayerRepo.
type MockRedisLayerRepoMockRecorder struct {
	mock *MockRedisLayerRepo
}

// NewMockRedisLayerRepo creates a new mock instance.
func NewMockRedisLayerRepo(ctrl *gomock.Controller) *MockRedisLayerRepo {
	mock := &MockRedisLayerRepo{ctrl: ctrl}
	mock.recorder = &MockRedisLayerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisLayerRepo) EXPECT() *MockRedisLayerRepoMockRecorder {
	return m.recorder
}

// Decr mocks base method.
func (m *MockRedisLayerRepo) Decr(key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decr indicates an expected call of Decr.
func (mr *MockRedisLayerRepoMockRecorder) Decr(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockRedisLayerRepo)(nil).Decr), key)
}

// DeleteKey mocks base method.
func (m *MockRedisLayerRepo) DeleteKey(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockRedisLayerRepoMockRecorder) DeleteKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockRedisLayerRepo)(nil).DeleteKey), key)
}

// GetDriver mocks base method.
func (m *MockRedisLayerRepo) GetDriver() kind.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriver")
	ret0, _ := ret[0].(kind.Driver)
	return ret0
}

// GetDriver indicates an expected call of GetDriver.
func (mr *MockRedisLayerRepoMockRecorder) GetDriver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriver", reflect.TypeOf((*MockRedisLayerRepo)(nil).GetDriver))
}

// GetVal mocks base method.
func (m *MockRedisLayerRepo) GetVal(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVal", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVal indicates an expected call of GetVal.
func (mr *MockRedisLayerRepoMockRecorder) GetVal(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVal", reflect.TypeOf((*MockRedisLayerRepo)(nil).GetVal), key)
}

// Incr mocks base method.
func (m *MockRedisLayerRepo) Incr(key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Incr indicates an expected call of Incr.
func (mr *MockRedisLayerRepoMockRecorder) Incr(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockRedisLayerRepo)(nil).Incr), key)
}

// Ping mocks base method.
func (m *MockRedisLayerRepo) Ping() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockRedisLayerRepoMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRedisLayerRepo)(nil).Ping))
}

// WriteKeyVal mocks base method.
func (m *MockRedisLayerRepo) WriteKeyVal(key, val string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyVal", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyVal indicates an expected call of WriteKeyVal.
func (mr *MockRedisLayerRepoMockRecorder) WriteKeyVal(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyVal", reflect.TypeOf((*MockRedisLayerRepo)(nil).WriteKeyVal), key, val)
}

// WriteKeyValAsJSON mocks base method.
func (m *MockRedisLayerRepo) WriteKeyValAsJSON(key string, val any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyValAsJSON", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyValAsJSON indicates an expected call of WriteKeyValAsJSON.
func (mr *MockRedisLayerRepoMockRecorder) WriteKeyValAsJSON(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyValAsJSON", reflect.TypeOf((*MockRedisLayerRepo)(nil).WriteKeyValAsJSON), key, val)
}

// WriteKeyValAsJSONTTL mocks base method.
func (m *MockRedisLayerRepo) WriteKeyValAsJSONTTL(key string, val any, insec int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyValAsJSONTTL", key, val, insec)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyValAsJSONTTL indicates an expected call of WriteKeyValAsJSONTTL.
func (mr *MockRedisLayerRepoMockRecorder) WriteKeyValAsJSONTTL(key, val, insec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyValAsJSONTTL", reflect.TypeOf((*MockRedisLayerRepo)(nil).WriteKeyValAsJSONTTL), key, val, insec)
}

// WriteKeyValTTL mocks base method.
func (m *MockRedisLayerRepo) WriteKeyValTTL(key, val string, insec int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteKeyValTTL", key, val, insec)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteKeyValTTL indicates an expected call of WriteKeyValTTL.
func (mr *MockRedisLayerRepoMockRecorder) WriteKeyValTTL(key, val, insec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteKeyValTTL", reflect.TypeOf((*MockRedisLayerRepo)(nil).WriteKeyValTTL), key, val, insec)
}

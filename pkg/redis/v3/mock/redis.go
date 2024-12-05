// Code generated by MockGen. DO NOT EDIT.
// Source: redis.go
//
// Generated by this command:
//
//	mockgen -source=redis.go -package=mock -destination=./mock/redis.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	redis0 "github.com/redis/go-redis/v9"
	gomock "go.uber.org/mock/gomock"

	health "github.com/bucketeer-io/bucketeer/pkg/health"
	redis "github.com/bucketeer-io/bucketeer/pkg/redis"
	v3 "github.com/bucketeer-io/bucketeer/pkg/redis/v3"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockClient) Check(arg0 context.Context) health.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(health.Status)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockClientMockRecorder) Check(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockClient)(nil).Check), arg0)
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Del mocks base method.
func (m *MockClient) Del(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockClientMockRecorder) Del(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockClient)(nil).Del), key)
}

// Dump mocks base method.
func (m *MockClient) Dump(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dump", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dump indicates an expected call of Dump.
func (mr *MockClientMockRecorder) Dump(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dump", reflect.TypeOf((*MockClient)(nil).Dump), key)
}

// Eval mocks base method.
func (m *MockClient) Eval(ctx context.Context, script string, keys []string, args ...any) *redis0.Cmd {
	m.ctrl.T.Helper()
	varargs := []any{ctx, script, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Eval", varargs...)
	ret0, _ := ret[0].(*redis0.Cmd)
	return ret0
}

// Eval indicates an expected call of Eval.
func (mr *MockClientMockRecorder) Eval(ctx, script, keys any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, script, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockClient)(nil).Eval), varargs...)
}

// Exists mocks base method.
func (m *MockClient) Exists(key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockClientMockRecorder) Exists(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockClient)(nil).Exists), key)
}

// Expire mocks base method.
func (m *MockClient) Expire(key string, expiration time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", key, expiration)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expire indicates an expected call of Expire.
func (mr *MockClientMockRecorder) Expire(key, expiration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockClient)(nil).Expire), key, expiration)
}

// Get mocks base method.
func (m *MockClient) Get(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), key)
}

// GetMulti mocks base method.
func (m *MockClient) GetMulti(keys []string, ignoreNotFound bool) ([]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMulti", keys, ignoreNotFound)
	ret0, _ := ret[0].([]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockClientMockRecorder) GetMulti(keys, ignoreNotFound any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockClient)(nil).GetMulti), keys, ignoreNotFound)
}

// Incr mocks base method.
func (m *MockClient) Incr(key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Incr indicates an expected call of Incr.
func (mr *MockClientMockRecorder) Incr(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockClient)(nil).Incr), key)
}

// IncrByFloat mocks base method.
func (m *MockClient) IncrByFloat(key string, value float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrByFloat", key, value)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrByFloat indicates an expected call of IncrByFloat.
func (mr *MockClientMockRecorder) IncrByFloat(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrByFloat", reflect.TypeOf((*MockClient)(nil).IncrByFloat), key, value)
}

// PFAdd mocks base method.
func (m *MockClient) PFAdd(key string, els ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range els {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFAdd", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PFAdd indicates an expected call of PFAdd.
func (mr *MockClientMockRecorder) PFAdd(key any, els ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, els...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFAdd", reflect.TypeOf((*MockClient)(nil).PFAdd), varargs...)
}

// PFCount mocks base method.
func (m *MockClient) PFCount(keys ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFCount", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PFCount indicates an expected call of PFCount.
func (mr *MockClientMockRecorder) PFCount(keys ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFCount", reflect.TypeOf((*MockClient)(nil).PFCount), keys...)
}

// PFMerge mocks base method.
func (m *MockClient) PFMerge(dest string, keys ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{dest}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFMerge", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PFMerge indicates an expected call of PFMerge.
func (mr *MockClientMockRecorder) PFMerge(dest any, keys ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{dest}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFMerge", reflect.TypeOf((*MockClient)(nil).PFMerge), varargs...)
}

// Pipeline mocks base method.
func (m *MockClient) Pipeline(tx bool) v3.PipeClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipeline", tx)
	ret0, _ := ret[0].(v3.PipeClient)
	return ret0
}

// Pipeline indicates an expected call of Pipeline.
func (mr *MockClientMockRecorder) Pipeline(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipeline", reflect.TypeOf((*MockClient)(nil).Pipeline), tx)
}

// Restore mocks base method.
func (m *MockClient) Restore(key string, ttl int64, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", key, ttl, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockClientMockRecorder) Restore(key, ttl, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockClient)(nil).Restore), key, ttl, value)
}

// Scan mocks base method.
func (m *MockClient) Scan(cursor uint64, key string, count int64) (uint64, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", cursor, key, count)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Scan indicates an expected call of Scan.
func (mr *MockClientMockRecorder) Scan(cursor, key, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockClient)(nil).Scan), cursor, key, count)
}

// Set mocks base method.
func (m *MockClient) Set(key string, val any, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, val, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockClientMockRecorder) Set(key, val, expiration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockClient)(nil).Set), key, val, expiration)
}

// SetNX mocks base method.
func (m *MockClient) SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, value, expiration)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX.
func (mr *MockClientMockRecorder) SetNX(ctx, key, value, expiration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockClient)(nil).SetNX), ctx, key, value, expiration)
}

// Stats mocks base method.
func (m *MockClient) Stats() redis.PoolStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(redis.PoolStats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockClientMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockClient)(nil).Stats))
}

// MockPipeClient is a mock of PipeClient interface.
type MockPipeClient struct {
	ctrl     *gomock.Controller
	recorder *MockPipeClientMockRecorder
}

// MockPipeClientMockRecorder is the mock recorder for MockPipeClient.
type MockPipeClientMockRecorder struct {
	mock *MockPipeClient
}

// NewMockPipeClient creates a new mock instance.
func NewMockPipeClient(ctrl *gomock.Controller) *MockPipeClient {
	mock := &MockPipeClient{ctrl: ctrl}
	mock.recorder = &MockPipeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipeClient) EXPECT() *MockPipeClientMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockPipeClient) Del(keys string) *redis0.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", keys)
	ret0, _ := ret[0].(*redis0.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockPipeClientMockRecorder) Del(keys any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockPipeClient)(nil).Del), keys)
}

// Exec mocks base method.
func (m *MockPipeClient) Exec() ([]redis0.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec")
	ret0, _ := ret[0].([]redis0.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPipeClientMockRecorder) Exec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPipeClient)(nil).Exec))
}

// Get mocks base method.
func (m *MockPipeClient) Get(key string) *redis0.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*redis0.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockPipeClientMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPipeClient)(nil).Get), key)
}

// Incr mocks base method.
func (m *MockPipeClient) Incr(key string) *redis0.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", key)
	ret0, _ := ret[0].(*redis0.IntCmd)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockPipeClientMockRecorder) Incr(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockPipeClient)(nil).Incr), key)
}

// PFAdd mocks base method.
func (m *MockPipeClient) PFAdd(key string, els ...string) *redis0.IntCmd {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range els {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFAdd", varargs...)
	ret0, _ := ret[0].(*redis0.IntCmd)
	return ret0
}

// PFAdd indicates an expected call of PFAdd.
func (mr *MockPipeClientMockRecorder) PFAdd(key any, els ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, els...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFAdd", reflect.TypeOf((*MockPipeClient)(nil).PFAdd), varargs...)
}

// PFCount mocks base method.
func (m *MockPipeClient) PFCount(keys ...string) *redis0.IntCmd {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFCount", varargs...)
	ret0, _ := ret[0].(*redis0.IntCmd)
	return ret0
}

// PFCount indicates an expected call of PFCount.
func (mr *MockPipeClientMockRecorder) PFCount(keys ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFCount", reflect.TypeOf((*MockPipeClient)(nil).PFCount), keys...)
}

// TTL mocks base method.
func (m *MockPipeClient) TTL(key string) *redis0.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", key)
	ret0, _ := ret[0].(*redis0.DurationCmd)
	return ret0
}

// TTL indicates an expected call of TTL.
func (mr *MockPipeClientMockRecorder) TTL(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockPipeClient)(nil).TTL), key)
}

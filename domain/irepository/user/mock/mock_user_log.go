// Code generated by MockGen. DO NOT EDIT.
// Source: ./Iuser_log_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/jettjia/go-ddd/domain/entity/user"
)

// MockIUserLogRepository is a mock of IUserLogRepository interface.
type MockIUserLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserLogRepositoryMockRecorder
}

// MockIUserLogRepositoryMockRecorder is the mock recorder for MockIUserLogRepository.
type MockIUserLogRepositoryMockRecorder struct {
	mock *MockIUserLogRepository
}

// NewMockIUserLogRepository creates a new mock instance.
func NewMockIUserLogRepository(ctrl *gomock.Controller) *MockIUserLogRepository {
	mock := &MockIUserLogRepository{ctrl: ctrl}
	mock.recorder = &MockIUserLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserLogRepository) EXPECT() *MockIUserLogRepositoryMockRecorder {
	return m.recorder
}

// SaveLog mocks base method.
func (m *MockIUserLogRepository) SaveLog(ctx context.Context, log *entity.UserLog) (*entity.UserLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLog", ctx, log)
	ret0, _ := ret[0].(*entity.UserLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveLog indicates an expected call of SaveLog.
func (mr *MockIUserLogRepositoryMockRecorder) SaveLog(ctx, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLog", reflect.TypeOf((*MockIUserLogRepository)(nil).SaveLog), ctx, log)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: event_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/paralleltree/go-leaderboard/internal/model"
)

// MockEventRepository is a mock of EventRepository interface.
type MockEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryMockRecorder
}

// MockEventRepositoryMockRecorder is the mock recorder for MockEventRepository.
type MockEventRepositoryMockRecorder struct {
	mock *MockEventRepository
}

// NewMockEventRepository creates a new mock instance.
func NewMockEventRepository(ctrl *gomock.Controller) *MockEventRepository {
	mock := &MockEventRepository{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepository) EXPECT() *MockEventRepositoryMockRecorder {
	return m.recorder
}

// GetEvent mocks base method.
func (m *MockEventRepository) GetEvent(ctx context.Context, id string) (model.Event, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", ctx, id)
	ret0, _ := ret[0].(model.Event)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockEventRepositoryMockRecorder) GetEvent(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockEventRepository)(nil).GetEvent), ctx, id)
}

// GetEvents mocks base method.
func (m *MockEventRepository) GetEvents(ctx context.Context, page, count int64) ([]model.Record[model.Event], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", ctx, page, count)
	ret0, _ := ret[0].([]model.Record[model.Event])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockEventRepositoryMockRecorder) GetEvents(ctx, page, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockEventRepository)(nil).GetEvents), ctx, page, count)
}

// RegisterEvent mocks base method.
func (m *MockEventRepository) RegisterEvent(ctx context.Context, event model.Event) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterEvent", ctx, event)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterEvent indicates an expected call of RegisterEvent.
func (mr *MockEventRepositoryMockRecorder) RegisterEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEvent", reflect.TypeOf((*MockEventRepository)(nil).RegisterEvent), ctx, event)
}

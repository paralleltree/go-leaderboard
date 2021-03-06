// Code generated by MockGen. DO NOT EDIT.
// Source: get_events_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/paralleltree/go-leaderboard/internal/model"
)

// MockGetEventsUsecase is a mock of GetEventsUsecase interface.
type MockGetEventsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockGetEventsUsecaseMockRecorder
}

// MockGetEventsUsecaseMockRecorder is the mock recorder for MockGetEventsUsecase.
type MockGetEventsUsecaseMockRecorder struct {
	mock *MockGetEventsUsecase
}

// NewMockGetEventsUsecase creates a new mock instance.
func NewMockGetEventsUsecase(ctrl *gomock.Controller) *MockGetEventsUsecase {
	mock := &MockGetEventsUsecase{ctrl: ctrl}
	mock.recorder = &MockGetEventsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetEventsUsecase) EXPECT() *MockGetEventsUsecaseMockRecorder {
	return m.recorder
}

// GetEvents mocks base method.
func (m *MockGetEventsUsecase) GetEvents(ctx context.Context, page, count int64) ([]model.Record[model.Event], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", ctx, page, count)
	ret0, _ := ret[0].([]model.Record[model.Event])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockGetEventsUsecaseMockRecorder) GetEvents(ctx, page, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockGetEventsUsecase)(nil).GetEvents), ctx, page, count)
}

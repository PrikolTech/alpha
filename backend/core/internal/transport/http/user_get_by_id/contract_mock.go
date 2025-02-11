// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go
//
// Generated by this command:
//
//	mockgen -package user_get_by_id_handler -source contract.go -destination contract_mock.go
//

// Package user_get_by_id_handler is a generated GoMock package.
package user_get_by_id

import (
	context "context"
	reflect "reflect"

	domain "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockuserUsecase is a mock of userUsecase interface.
type MockuserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockuserUsecaseMockRecorder
	isgomock struct{}
}

// MockuserUsecaseMockRecorder is the mock recorder for MockuserUsecase.
type MockuserUsecaseMockRecorder struct {
	mock *MockuserUsecase
}

// NewMockuserUsecase creates a new mock instance.
func NewMockuserUsecase(ctrl *gomock.Controller) *MockuserUsecase {
	mock := &MockuserUsecase{ctrl: ctrl}
	mock.recorder = &MockuserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserUsecase) EXPECT() *MockuserUsecaseMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockuserUsecase) Handle(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockuserUsecaseMockRecorder) Handle(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockuserUsecase)(nil).Handle), ctx, id)
}

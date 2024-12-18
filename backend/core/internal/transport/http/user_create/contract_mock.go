// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go
//
// Generated by this command:
//
//	mockgen -package user_create_handler -source contract.go -destination contract_mock.go
//

// Package user_create_handler is a generated GoMock package.
package user_create_handler

import (
	context "context"
	reflect "reflect"

	domain "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
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
func (m *MockuserUsecase) Handle(ctx context.Context, in domain.UserCreateIn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockuserUsecaseMockRecorder) Handle(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockuserUsecase)(nil).Handle), ctx, in)
}

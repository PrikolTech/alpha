// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go
//
// Generated by this command:
//
//	mockgen -package user_list -source contract.go -destination contract_mock.go
//

// Package user_list is a generated GoMock package.
package user_list

import (
	context "context"
	reflect "reflect"

	domain "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	gomock "go.uber.org/mock/gomock"
)

// Mockusecase is a mock of usecase interface.
type Mockusecase struct {
	ctrl     *gomock.Controller
	recorder *MockusecaseMockRecorder
	isgomock struct{}
}

// MockusecaseMockRecorder is the mock recorder for Mockusecase.
type MockusecaseMockRecorder struct {
	mock *Mockusecase
}

// NewMockusecase creates a new mock instance.
func NewMockusecase(ctrl *gomock.Controller) *Mockusecase {
	mock := &Mockusecase{ctrl: ctrl}
	mock.recorder = &MockusecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockusecase) EXPECT() *MockusecaseMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *Mockusecase) Handle(ctx context.Context, in domain.UserListIn) (*domain.UserListOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, in)
	ret0, _ := ret[0].(*domain.UserListOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockusecaseMockRecorder) Handle(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*Mockusecase)(nil).Handle), ctx, in)
}

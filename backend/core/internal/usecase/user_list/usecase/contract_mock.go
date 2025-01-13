// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go
//
// Generated by this command:
//
//	mockgen -package usecase -source contract.go -destination contract_mock.go
//

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	domain "github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockuserRepo is a mock of userRepo interface.
type MockuserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepoMockRecorder
	isgomock struct{}
}

// MockuserRepoMockRecorder is the mock recorder for MockuserRepo.
type MockuserRepoMockRecorder struct {
	mock *MockuserRepo
}

// NewMockuserRepo creates a new mock instance.
func NewMockuserRepo(ctrl *gomock.Controller) *MockuserRepo {
	mock := &MockuserRepo{ctrl: ctrl}
	mock.recorder = &MockuserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserRepo) EXPECT() *MockuserRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockuserRepo) Get(ctx context.Context, in domain.UserListIn) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, in)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockuserRepoMockRecorder) Get(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockuserRepo)(nil).Get), ctx, in)
}

// GetTotalCount mocks base method.
func (m *MockuserRepo) GetTotalCount(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalCount", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalCount indicates an expected call of GetTotalCount.
func (mr *MockuserRepoMockRecorder) GetTotalCount(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalCount", reflect.TypeOf((*MockuserRepo)(nil).GetTotalCount), ctx)
}

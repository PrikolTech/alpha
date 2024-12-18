// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// UserCreate implements userCreate operation.
//
// Создать нового пользователя.
//
// POST /v1/users
func (UnimplementedHandler) UserCreate(ctx context.Context, req *UserCreateRequest) (r UserCreateRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UserGetAll implements userGetAll operation.
//
// Получить всех пользователей.
//
// GET /v1/users
func (UnimplementedHandler) UserGetAll(ctx context.Context, params UserGetAllParams) (r *UserGetAllResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UserGetById implements userGetById operation.
//
// Получить пользователя по id.
//
// GET /v1/users/{id}
func (UnimplementedHandler) UserGetById(ctx context.Context, params UserGetByIdParams) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

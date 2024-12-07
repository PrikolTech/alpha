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
func (UnimplementedHandler) UserCreate(ctx context.Context, req *UserCreateRequest) error {
	return ht.ErrNotImplemented
}

// UserGetById implements userGetById operation.
//
// Получить пользователя по id.
//
// GET /v1/users/{id}
func (UnimplementedHandler) UserGetById(ctx context.Context, params UserGetByIdParams) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
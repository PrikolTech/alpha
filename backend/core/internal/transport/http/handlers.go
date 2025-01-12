package http

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	user_create_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_create"
	user_get_by_id_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_get_by_id"
	user_list_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_list"
)

type Handlers struct {
	UserCreate  *user_create_handler.Handler
	UserGetById *user_get_by_id_handler.Handler
	UserList    *user_list_handler.Handler
}

func (m mux) UserCreate(ctx context.Context, req *api.UserCreateRequest) (api.UserCreateRes, error) {
	return m.handlers.UserCreate.Handle(ctx, req)
}

func (m mux) UserGetById(ctx context.Context, params api.UserGetByIdParams) (api.UserGetByIdRes, error) {
	return m.handlers.UserGetById.Handle(ctx, params)
}

func (h mux) UserList(ctx context.Context, params api.UserListParams) (api.UserListRes, error) {
	return h.handlers.UserList.Handle(ctx, params)
}

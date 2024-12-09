//go:generate ogen --config ../../../.ogen.yml --target ../../generated/api  --package api --clean ../../../../../docs/api/openapi.yml

package http

import (
	"context"

	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
	user_create_handler "github.com/PrikolTech/alpha/backend/core/internal/transport/http/user_create"
)

type Handlers struct {
	UserCreate *user_create_handler.Handler
}

type mux struct {
	*api.UnimplementedHandler
	handlers Handlers
}

func New(handlers Handlers) *api.Server {
	mux := mux{
		UnimplementedHandler: new(api.UnimplementedHandler),
		handlers:             handlers,
	}

	server, err := api.NewServer(mux)
	if err != nil {
		return nil
	}

	return server
}

func (h mux) UserCreate(ctx context.Context, req *api.UserCreateRequest) (api.UserCreateRes, error) {
	return h.handlers.UserCreate.Handle(ctx, req)
}

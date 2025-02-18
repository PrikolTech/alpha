//go:generate go tool ogen --config ../../../.ogen.yml --target ../../generated/api  --package api --clean ../../../../../docs/api/openapi.yml

package http

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-faster/jx"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"

	"github.com/PrikolTech/alpha/backend/core/internal/generated/api"
)

type mux struct {
	*api.UnimplementedHandler
	handlers Handlers
	logger   *slog.Logger
}

func New(logger *slog.Logger, handlers Handlers) *api.Server {
	mux := mux{
		UnimplementedHandler: new(api.UnimplementedHandler),
		handlers:             handlers,
		logger:               logger,
	}

	server, err := api.NewServer(mux, api.WithErrorHandler(mux.handleError))
	if err != nil {
		return nil
	}

	return server
}

func (m mux) handleError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var ogenErr ogenerrors.Error
	switch {
	case errors.As(err, &ogenErr):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(ogenErr.Code())
		e := jx.GetEncoder()
		e.ObjStart()
		e.FieldStart("message")
		e.StrEscape(err.Error())
		e.ObjEnd()
		_, _ = w.Write(e.Bytes())

	case errors.Is(err, ht.ErrNotImplemented):
		w.WriteHeader(http.StatusNotImplemented)

	case errors.Is(err, context.Canceled):
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusInternalServerError)
		m.logger.Error("internal server error", slog.String("err", err.Error()))
	}
}

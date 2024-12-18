// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

type codeRecorder struct {
	http.ResponseWriter
	status int
}

func (c *codeRecorder) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func recordError(string, error) {}

// handleUserCreateRequest handles userCreate operation.
//
// Создать нового пользователя.
//
// POST /v1/users
func (s *Server) handleUserCreateRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UserCreateOperation,
			ID:   "userCreate",
		}
	)
	request, close, err := s.decodeUserCreateRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response UserCreateRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UserCreateOperation,
			OperationSummary: "Создать нового пользователя",
			OperationID:      "userCreate",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *UserCreateRequest
			Params   = struct{}
			Response = UserCreateRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UserCreate(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.UserCreate(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserCreateResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUserGetAllRequest handles userGetAll operation.
//
// Получить всех пользователей.
//
// GET /v1/users
func (s *Server) handleUserGetAllRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UserGetAllOperation,
			ID:   "userGetAll",
		}
	)
	params, err := decodeUserGetAllParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *UserGetAllResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UserGetAllOperation,
			OperationSummary: "Получить всех пользователей",
			OperationID:      "userGetAll",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "per",
					In:   "query",
				}: params.Per,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UserGetAllParams
			Response = *UserGetAllResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUserGetAllParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UserGetAll(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UserGetAll(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserGetAllResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUserGetByIdRequest handles userGetById operation.
//
// Получить пользователя по id.
//
// GET /v1/users/{id}
func (s *Server) handleUserGetByIdRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: UserGetByIdOperation,
			ID:   "userGetById",
		}
	)
	params, err := decodeUserGetByIdParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *User
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    UserGetByIdOperation,
			OperationSummary: "Получить пользователя по id",
			OperationID:      "userGetById",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = UserGetByIdParams
			Response = *User
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUserGetByIdParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UserGetById(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UserGetById(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserGetByIdResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

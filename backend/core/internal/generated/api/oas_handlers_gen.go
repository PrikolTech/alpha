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

// handleProjectCreateRequest handles projectCreate operation.
//
// Создать новый проект.
//
// POST /v1/projects
func (s *Server) handleProjectCreateRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ProjectCreateOperation,
			ID:   "projectCreate",
		}
	)
	request, close, err := s.decodeProjectCreateRequest(r)
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

	var response *ProjectCreateCreated
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ProjectCreateOperation,
			OperationSummary: "Создать новый проект",
			OperationID:      "projectCreate",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *ProjectCreateRequest
			Params   = struct{}
			Response = *ProjectCreateCreated
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
				err = s.h.ProjectCreate(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.ProjectCreate(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeProjectCreateResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleProjectDeleteByIdRequest handles projectDeleteById operation.
//
// Удалить проект по id.
//
// DELETE /v1/projects/{id}
func (s *Server) handleProjectDeleteByIdRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ProjectDeleteByIdOperation,
			ID:   "projectDeleteById",
		}
	)
	params, err := decodeProjectDeleteByIdParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ProjectDeleteByIdNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ProjectDeleteByIdOperation,
			OperationSummary: "Удалить проект по id",
			OperationID:      "projectDeleteById",
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
			Params   = ProjectDeleteByIdParams
			Response = *ProjectDeleteByIdNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackProjectDeleteByIdParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ProjectDeleteById(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.ProjectDeleteById(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeProjectDeleteByIdResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleProjectGetAllRequest handles projectGetAll operation.
//
// Получить все проекты с пагинацией.
//
// GET /v1/projects
func (s *Server) handleProjectGetAllRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ProjectGetAllOperation,
			ID:   "projectGetAll",
		}
	)
	params, err := decodeProjectGetAllParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ProjectGetAllResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ProjectGetAllOperation,
			OperationSummary: "Получить все проекты с пагинацией",
			OperationID:      "projectGetAll",
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
			Params   = ProjectGetAllParams
			Response = *ProjectGetAllResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackProjectGetAllParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ProjectGetAll(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ProjectGetAll(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeProjectGetAllResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleProjectGetByIdRequest handles projectGetById operation.
//
// Получить проект по id.
//
// GET /v1/projects/{id}
func (s *Server) handleProjectGetByIdRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	ctx := r.Context()

	var (
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: ProjectGetByIdOperation,
			ID:   "projectGetById",
		}
	)
	params, err := decodeProjectGetByIdParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Project
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    ProjectGetByIdOperation,
			OperationSummary: "Получить проект по id",
			OperationID:      "projectGetById",
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
			Params   = ProjectGetByIdParams
			Response = *Project
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackProjectGetByIdParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ProjectGetById(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ProjectGetById(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeProjectGetByIdResponse(response, w); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

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

	var response UserGetByIdRes
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
			Response = UserGetByIdRes
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

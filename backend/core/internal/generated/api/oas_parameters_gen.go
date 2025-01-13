// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// ProjectDeleteByIdParams is parameters of projectDeleteById operation.
type ProjectDeleteByIdParams struct {
	// Id проекта.
	ID uuid.UUID
}

func unpackProjectDeleteByIdParams(packed middleware.Parameters) (params ProjectDeleteByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeProjectDeleteByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params ProjectDeleteByIdParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ProjectGetAllParams is parameters of projectGetAll operation.
type ProjectGetAllParams struct {
	// Номер страницы.
	Page OptInt
	// Количество записей на странице.
	PerPage OptInt
}

func unpackProjectGetAllParams(packed middleware.Parameters) (params ProjectGetAllParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "perPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PerPage = v.(OptInt)
		}
	}
	return params
}

func decodeProjectGetAllParams(args [0]string, argsEscaped bool, r *http.Request) (params ProjectGetAllParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(1)
		params.Page.SetTo(val)
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: perPage.
	{
		val := int(20)
		params.PerPage.SetTo(val)
	}
	// Decode query: perPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "perPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PerPage.SetTo(paramsDotPerPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "perPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ProjectGetByIdParams is parameters of projectGetById operation.
type ProjectGetByIdParams struct {
	// Id проекта.
	ID uuid.UUID
}

func unpackProjectGetByIdParams(packed middleware.Parameters) (params ProjectGetByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeProjectGetByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params ProjectGetByIdParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UserGetByIdParams is parameters of userGetById operation.
type UserGetByIdParams struct {
	// Id пользователя.
	ID uuid.UUID
}

func unpackUserGetByIdParams(packed middleware.Parameters) (params UserGetByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeUserGetByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params UserGetByIdParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UserListParams is parameters of userList operation.
type UserListParams struct {
	// Номер страницы.
	Page OptInt
	// Количество записей на странице.
	PerPage OptInt
	// Сортировка списка.
	Sorting OptSorting
	// Почтовый адрес.
	Email OptString
	// Имя.
	FirstName OptString
	// Отчество.
	MiddleName OptString
	// Фамилия.
	LastName OptString
	// Дата и время создания.
	CreatedAt OptDateTimeFilter
	// Дата и время обновления.
	UpdatedAt OptDateTimeFilter
}

func unpackUserListParams(packed middleware.Parameters) (params UserListParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "perPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PerPage = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sorting",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sorting = v.(OptSorting)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "email",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Email = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "firstName",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FirstName = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "middleName",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.MiddleName = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "lastName",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.LastName = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "createdAt",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CreatedAt = v.(OptDateTimeFilter)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "updatedAt",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UpdatedAt = v.(OptDateTimeFilter)
		}
	}
	return params
}

func decodeUserListParams(args [0]string, argsEscaped bool, r *http.Request) (params UserListParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(1)
		params.Page.SetTo(val)
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: perPage.
	{
		val := int(20)
		params.PerPage.SetTo(val)
	}
	// Decode query: perPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "perPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PerPage.SetTo(paramsDotPerPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "perPage",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sorting.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sorting",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "field", Required: true}, {Name: "direction", Required: true}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortingVal Sorting
				if err := func() error {
					return paramsDotSortingVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Sorting.SetTo(paramsDotSortingVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Sorting.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sorting",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: email.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "email",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEmailVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEmailVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Email.SetTo(paramsDotEmailVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "email",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: firstName.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "firstName",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFirstNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotFirstNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.FirstName.SetTo(paramsDotFirstNameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "firstName",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: middleName.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "middleName",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotMiddleNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotMiddleNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.MiddleName.SetTo(paramsDotMiddleNameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "middleName",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: lastName.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "lastName",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLastNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotLastNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.LastName.SetTo(paramsDotLastNameVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "lastName",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: createdAt.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "createdAt",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "start", Required: false}, {Name: "end", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCreatedAtVal DateTimeFilter
				if err := func() error {
					return paramsDotCreatedAtVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.CreatedAt.SetTo(paramsDotCreatedAtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "createdAt",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: updatedAt.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "updatedAt",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "start", Required: false}, {Name: "end", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUpdatedAtVal DateTimeFilter
				if err := func() error {
					return paramsDotUpdatedAtVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.UpdatedAt.SetTo(paramsDotUpdatedAtVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "updatedAt",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

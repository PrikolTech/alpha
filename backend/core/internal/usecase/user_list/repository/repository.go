package repository

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, in domain.UserListIn) ([]domain.User, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("*").
		From("_user").
		Where(r.buildWhereQuery(in.Filters)).
		Limit(uint64(in.PerPage)).
		Offset(uint64((in.Page - 1) * in.PerPage))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var entities []entity

	err = r.db.SelectContext(ctx, &entities, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return lo.Map(entities, func(item entity, _ int) domain.User { return item.toDomain() }), nil
}

func (r *Repository) buildWhereQuery(filters domain.UserListFilters) sq.Sqlizer {
	var res sq.And

	if filters.Email != nil {
		res = append(res, r.buildLikeQuery("email", filters.Email))
	}

	if filters.FirstName != nil {
		res = append(res, r.buildLikeQuery("first_name", filters.FirstName))
	}

	if filters.MiddleName != nil {
		res = append(res, r.buildLikeQuery("middle_name", filters.MiddleName))
	}

	if filters.LastName != nil {
		res = append(res, r.buildLikeQuery("last_name", filters.LastName))
	}

	if filters.CreatedAt != nil {
		if filters.CreatedAt.Start != nil {
			res = append(res, sq.GtOrEq{"created_at": *filters.CreatedAt.Start})
		}
		if filters.CreatedAt.End != nil {
			res = append(res, sq.LtOrEq{"created_at": *filters.CreatedAt.End})
		}
	}

	return res
}

func (r *Repository) buildLikeQuery(field string, value *string) sq.ILike {
	return sq.ILike{field: fmt.Sprintf("%%%s%%", strings.ToLower(*value))}
}

func (r *Repository) GetTotalCount(ctx context.Context) (int, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("COUNT(*)").
		From("_user")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	var totalCount int
	err = r.db.GetContext(ctx, &totalCount, query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to exec query: %w", err)
	}

	return totalCount, nil
}

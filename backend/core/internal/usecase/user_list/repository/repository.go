package repository

import (
	"context"
	"fmt"

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

func (r *Repository) GetAll(ctx context.Context, in domain.UserListIn) ([]domain.User, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("*").
		From("_user").
		Limit(uint64(in.PerPage)).Offset(uint64((in.Page - 1) * in.PerPage))

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

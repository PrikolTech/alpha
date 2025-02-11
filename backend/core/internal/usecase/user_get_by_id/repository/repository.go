package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_by_id/domain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	op := "user get by id"
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("*").
		From("_user").
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build %q query: %w", op, err)
	}

	var entities []entity

	query = fmt.Sprintf("-- %s\n%s", query, op)
	err = r.db.SelectContext(ctx, &entities, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec %q query: %w", op, err)
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return entities[0].toDomain(), nil
}

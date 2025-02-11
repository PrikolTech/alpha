package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/project_create/domain"
)

type Repository struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func New(db *sqlx.DB, getter *trmsqlx.CtxGetter) *Repository {
	return &Repository{db: db, getter: getter}
}

func (r *Repository) ExistsByCode(ctx context.Context, code string) (bool, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id").
		From("project").
		Where(sq.Eq{"code": code})

	query, args, err := builder.ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.getter.DefaultTrOrDB(ctx, r.db).QueryContext(ctx, query, args...)
	if err != nil {
		return false, fmt.Errorf("failed to exec query: %w", err)
	}

	if rows.Next() {
		rows.Close()
		return true, nil
	}

	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("failed to iterate over rows: %w", err)
	}

	return false, nil
}

func (r *Repository) Create(ctx context.Context, in domain.ProjectCreateIn) error {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("project").
		Columns("name", "description", "code").
		Values(in.Name, in.Description, in.Code)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.getter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}

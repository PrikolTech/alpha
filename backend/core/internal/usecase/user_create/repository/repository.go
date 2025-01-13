package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_create/domain"
)

type Repository struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func New(db *sqlx.DB, getter *trmsqlx.CtxGetter) *Repository {
	return &Repository{db: db, getter: getter}
}

func (r *Repository) Create(ctx context.Context, in domain.UserCreateIn) error {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("_user").
		Columns("email", "first_name", "middle_name", "last_name").
		Values(in.Email, in.FirstName, in.MiddleName, in.LastName)

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

func (r *Repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id").
		From("_user").
		Where(sq.Eq{"email": email})

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

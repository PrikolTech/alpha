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
	op := "user create"
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("_user").
		Columns("email", "first_name", "middle_name", "last_name").
		Values(in.Email, in.FirstName, in.MiddleName, in.LastName)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build %q query: %w", op, err)
	}

	query = fmt.Sprintf("-- %s\n%s", query, op)
	_, err = r.getter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec %q query: %w", op, err)
	}

	return nil
}

func (r *Repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	op := "exists by email"
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id").
		From("_user").
		Where(sq.Eq{"email": email})

	query, args, err := builder.ToSql()
	if err != nil {
		return false, fmt.Errorf("build %q query: %w", op, err)
	}

	query = fmt.Sprintf("-- %s\n%s", query, op)
	rows, err := r.getter.DefaultTrOrDB(ctx, r.db).QueryContext(ctx, query, args...)
	if err != nil {
		return false, fmt.Errorf("exec %q query: %w", op, err)
	}

	if rows.Next() {
		rows.Close()
		return true, nil
	}

	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("iterate over rows: %w", err)
	}

	return false, nil
}

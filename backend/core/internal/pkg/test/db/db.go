package test_db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/PrikolTech/alpha/backend/core/internal/pkg/psql"
)

type Container struct {
	db      *sqlx.DB
	builder *sq.StatementBuilderType
}

func New(db *sqlx.DB, builder *sq.StatementBuilderType) *Container {
	return &Container{db, builder}
}

func (c *Container) DB() *sqlx.DB { return c.db }

func (c *Container) Close() error { return c.db.Close() }

func NewPsql() (*Container, error) {
	db, err := psql.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &Container{db: db, builder: &builder}, nil
}

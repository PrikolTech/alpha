package psql

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(opts ...OptionFunc) (*sqlx.DB, error) {
	cfg, err := pgx.ParseConfig("")
	if err != nil {
		return nil, err
	}

	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}

	db, err := sqlx.Open(options.driverName, stdlib.RegisterConnConfig(cfg))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

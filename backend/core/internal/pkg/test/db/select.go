package test_db

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/samber/lo"
)

func SelectEntitiesByIdCol[T, U any](c *Container, table string, idCol string, ids []U) ([]T, error) {
	var entity T
	fields := toFields(entity)

	builder := c.builder.
		Select(lo.Keys(fields)...).
		From(table).
		Where(sq.Eq{idCol: ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query for '%s': %w", table, err)
	}

	var entities []T
	err = c.db.Select(&entities, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query for '%s': %w", table, err)
	}

	return entities, nil
}

func SelectEntitiesById[T, U any](c *Container, table string, ids []U) ([]T, error) {
	return SelectEntitiesByIdCol[T, U](c, table, "id", ids)
}

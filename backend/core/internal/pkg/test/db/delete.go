package test_db

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteEntitiesByIdCol deletes entities from table by ids in column idCol.
func DeleteEntitiesByIdCol[U any](c *Container, table string, idCol string, ids []U) error {
	builder := c.builder.
		Delete(table).
		Where(sq.Eq{idCol: ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query for '%s': %w", table, err)
	}

	_, err = c.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query for '%s': %w", table, err)
	}

	return nil
}

// DeleteEntitiesById deletes entities from table by ids in column "id".
func DeleteEntitiesById[U any](c *Container, table string, ids []U) error {
	return DeleteEntitiesByIdCol(c, table, "id", ids)
}

// DeleteEntityByIdCol deletes entity from table by id in column idCol.
func DeleteEntityByIdCol[U any](c *Container, table string, idCol string, id U) error {
	return DeleteEntitiesByIdCol(c, table, idCol, []U{id})
}

// DeleteEntityById deletes entity from table by id in column "id".
func DeleteEntityById[U any](c *Container, table string, id U) error {
	return DeleteEntitiesById(c, table, []U{id})
}

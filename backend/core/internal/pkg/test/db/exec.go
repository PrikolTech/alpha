package test_db

import (
	"fmt"
	"slices"

	sq "github.com/Masterminds/squirrel"
	"github.com/samber/lo"
)

func insertEntities[T, U any](c *Container, table string, entities []T, idCol string, idSkip bool) ([]U, error) {
	if len(entities) == 0 {
		return make([]U, 0), nil
	}

	fieldsSlice := lo.Map(entities, func(item T, _ int) []field {
		fields := toFields(item)
		if idSkip {
			fields = slices.DeleteFunc(fields, func(field field) bool {
				return field.key == idCol
			})
		}
		return fields
	})

	keys := lo.Map(fieldsSlice[0], func(item field, _ int) string { return item.key })

	builder := c.builder.
		Insert(table).
		Columns(keys...)

	for _, fields := range fieldsSlice {
		values := lo.Map(fields, func(item field, _ int) any { return item.value })
		builder = builder.Values(values...)
	}

	builder = builder.Suffix("RETURNING " + idCol)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query for %q: %w", table, err)
	}

	var ids []U
	err = c.db.Select(&ids, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query for %q: %w", table, err)
	}

	return ids, nil
}

// InsertEntitiesById skips fields with tag "id", inserts entities into table and returns ids from column "id".
func InsertEntitiesById[T, U any](c *Container, table string, entities []T) ([]U, error) {
	return insertEntities[T, U](c, table, entities, "id", true)
}

// InsertEntitiesByIdCol skips fields with tag idCol, inserts entities into table and returns ids from column idCol.
func InsertEntitiesByIdCol[T, U any](c *Container, table string, entities []T, idCol string) ([]U, error) {
	return insertEntities[T, U](c, table, entities, idCol, true)
}

// InsertEntitiesWithId inserts entities into table and returns ids from column "id".
func InsertEntitiesWithId[T, U any](c *Container, table string, entities []T) ([]U, error) {
	return insertEntities[T, U](c, table, entities, "id", false)
}

// InsertEntitiesWithIdCol inserts entities into table and returns ids from column idCol.
func InsertEntitiesWithIdCol[T, U any](c *Container, table string, entities []T, idCol string) ([]U, error) {
	return insertEntities[T, U](c, table, entities, idCol, false)
}

func insertEntity[T, U any](c *Container, table string, entity T, idCol string, idSkip bool) (U, error) {
	ids, err := insertEntities[T, U](c, table, []T{entity}, idCol, idSkip)
	if err != nil {
		var id U
		return id, err
	}
	return ids[0], nil
}

// InsertEntityById skips field with tag "id", inserts entity into table and returns id from column "id".
func InsertEntityById[T, U any](c *Container, table string, entity T) (U, error) {
	return insertEntity[T, U](c, table, entity, "id", true)
}

// InsertEntityByIdCol skips field with tag idCol, inserts entity into table and returns id from column idCol.
func InsertEntityByIdCol[T, U any](c *Container, table string, entity T, idCol string) (U, error) {
	return insertEntity[T, U](c, table, entity, idCol, true)
}

// InsertEntityWithId inserts entity into table and returns id from column "id".
func InsertEntityWithId[T, U any](c *Container, table string, entity T) (U, error) {
	return insertEntity[T, U](c, table, entity, "id", false)
}

// InsertEntityWithIdCol inserts entity into table and returns id from column idCol.
func InsertEntityWithIdCol[T, U any](c *Container, table string, entity T, idCol string) (U, error) {
	return insertEntity[T, U](c, table, entity, idCol, false)
}

// DeleteEntitiesByIdCol deletes entities from table by ids in column idCol.
func DeleteEntitiesByIdCol[U any](c *Container, table string, idCol string, ids []U) error {
	builder := c.builder.
		Delete(table).
		Where(sq.Eq{idCol: ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query for %q: %w", table, err)
	}

	_, err = c.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("exec query for %q: %w", table, err)
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

// SelectEntitiesByIdCol selects entities from table by ids in column idCol
func SelectEntitiesByIdCol[T, U any](c *Container, table string, idCol string, ids []U) ([]T, error) {
	var entity T
	fields := toFields(entity)
	keys := lo.Map(fields, func(item field, _ int) string { return item.key })

	builder := c.builder.
		Select(keys...).
		From(table).
		Where(sq.Eq{idCol: ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query for %q: %w", table, err)
	}

	var entities []T
	err = c.db.Select(&entities, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query for %q: %w", table, err)
	}

	return entities, nil
}

// SelectEntitiesById selects entities from table by ids in column "id".
func SelectEntitiesById[T, U any](c *Container, table string, ids []U) ([]T, error) {
	return SelectEntitiesByIdCol[T, U](c, table, "id", ids)
}

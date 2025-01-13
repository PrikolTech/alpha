package test_db

import (
	"reflect"
)

type Field struct {
	Key   string
	Value any
}

func toFields[T any](item T) []Field {
	v := reflect.ValueOf(item)
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return nil
	}

	fields := make([]Field, v.NumField())
	for i := range fields {
		fields[i].Key = t.Field(i).Tag.Get("db")
		fields[i].Value = v.Field(i).Interface()
	}

	return fields
}

package test_db

import "reflect"

type Fields map[string]any

func toFields[T any](item T) Fields {
	v := reflect.ValueOf(item)
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return nil
	}

	fields := make(Fields, v.NumField())
	for i := range v.NumField() {
		key := t.Field(i).Tag.Get("db")
		fields[key] = v.Field(i).Interface()
	}

	return fields
}

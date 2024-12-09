package test_db

import "github.com/brianvoe/gofakeit/v7"

type GenerateOption[T any] func(entity *T)

func GenerateEntity[T any](opts ...GenerateOption[T]) (T, error) {
	var entity T

	if err := gofakeit.Struct(&entity); err != nil {
		return entity, nil
	}

	for _, opt := range opts {
		opt(&entity)
	}

	return entity, nil
}

func GenerateEntities[T any](n int, opts ...GenerateOption[T]) ([]T, error) {
	entities := make([]T, n)

	var err error
	for i := range entities {
		entities[i], err = GenerateEntity(opts...)
		if err != nil {
			return nil, err
		}
	}

	return entities, nil
}

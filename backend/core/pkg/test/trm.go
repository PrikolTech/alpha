package test

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type TrManager struct{}

func (m *TrManager) Do(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	return fn(ctx)
}

func (m *TrManager) DoWithSettings(ctx context.Context, _ trm.Settings, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func NewTrManager() *TrManager {
	return &TrManager{}
}

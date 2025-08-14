package mocks

import (
	"context"

	"github.com/juanMaAV92/go-utils/cache"
	"github.com/stretchr/testify/mock"
)

type Cache struct {
	mock.Mock
}

func (m *Cache) Set(ctx context.Context, key string, value interface{}, opts ...cache.SetOption) error {
	args := m.Called(ctx, key, value, opts)
	return args.Error(0)
}

func (m *Cache) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

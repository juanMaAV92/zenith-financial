package mocks

import (
	"context"

	"github.com/juanMaAV92/go-utils/log"
	"github.com/stretchr/testify/mock"
)

type Logger struct {
	mock.Mock
}

func (m *Logger) Fatal(ctx context.Context, step, message string, options ...log.Opts) {
	m.Called(ctx, step, message, options)
}

func (m *Logger) Error(ctx context.Context, step, message string, options ...log.Opts) {
	m.Called(ctx, step, message, options)
}

func (m *Logger) Warning(ctx context.Context, step, message string, options ...log.Opts) {
	m.Called(ctx, step, message, options)
}

func (m *Logger) Info(ctx context.Context, step, message string, options ...log.Opts) {
	m.Called(ctx, step, message, options)
}

func (m *Logger) Debug(ctx context.Context, step, message string, options ...log.Opts) {
	m.Called(ctx, step, message, options)
}

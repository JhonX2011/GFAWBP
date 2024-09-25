package databasemock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockGormClient struct {
	mock.Mock
}

func (m *MockGormClient) GetDB(ctx context.Context) *gorm.DB {
	args := m.Called(ctx)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormClient) RetryQuery(ctx context.Context, queryFunc func() *gorm.DB) *gorm.DB {
	args := m.Called(ctx, queryFunc)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormClient) Begin(ctx context.Context) (context.Context, func(), error) {
	args := m.Called(ctx)
	return args.Get(0).(context.Context), args.Get(1).(func()), args.Error(2)
}

func (m *MockGormClient) Commit(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockGormClient) Rollback(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

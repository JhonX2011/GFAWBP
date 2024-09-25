package databasemock

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockConnections struct {
	mock.Mock
}

func (m *MockConnections) Get(name string) (*sql.DB, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sql.DB), args.Error(1)
}

func (m *MockConnections) List() []*sql.DB {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*sql.DB)
	}
	return make([]*sql.DB, 0)
}

func (m *MockConnections) Close() error {
	args := m.Called()
	if args.Error(0) != nil {
		return args.Error(0)
	}
	return nil
}

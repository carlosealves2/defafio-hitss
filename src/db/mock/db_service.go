package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
)

type MockDbService struct {
	mock.Mock
}

func (m *MockDbService) CreateUser(ctx context.Context, arg postgres.CreateUserParams) (postgres.User, error) {
	return postgres.User{}, nil
}

func (m *MockDbService) DeleteUser(ctx context.Context, id int64) (int64, error) {
	return 0, nil
}

func (m *MockDbService) GetUser(ctx context.Context, id int64) (postgres.User, error) {
	return postgres.User{}, nil
}

func (m *MockDbService) ListUsers(ctx context.Context) ([]postgres.User, error) {
	return nil, nil
}

func (m *MockDbService) UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.User, error) {
	return postgres.User{}, nil
}

func NewMockDbService() *MockDbService {
	return &MockDbService{}
}

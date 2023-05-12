package protocols

import (
	"context"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
)

type IDBService interface {
	CreateUser(ctx context.Context, arg postgres.CreateUserParams) (postgres.User, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
	GetUser(ctx context.Context, id int64) (postgres.User, error)
	ListUsers(ctx context.Context) ([]postgres.User, error)
	UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.User, error)
}

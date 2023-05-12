package protocols

import (
	"context"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
)

type ICreateUserService interface {
	Create(ctx context.Context, user postgres.CreateUserParams) error
	Validate(user postgres.CreateUserParams) error
	ObfuscateInformation(ctx context.Context, user postgres.CreateUserParams, fields []string) (*postgres.CreateUserParams, error)
}

package usecases

import (
	"context"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
	"github.com/suportebeloj/desafio-hitss/src/utils/cerrors"
	"github.com/suportebeloj/desafio-hitss/src/utils/validators"
)

type CreateUserService struct {
	userServices *postgres.Queries
}

func NewProcessUserData() *CreateUserService {
	return &CreateUserService{}
}

func (p *CreateUserService) Create(ctx context.Context, user postgres.CreateUserParams) error {
	if err := p.validate(user); err != nil {
		return err
	}

	_, err := p.userServices.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (p *CreateUserService) validate(user postgres.CreateUserParams) error {

	if (user.Name == "" || user.Surname == "") || (validators.InvalidChar(user.Name) || validators.InvalidChar(user.Surname)) {
		return &cerrors.InvalidNameOrSurnameError{}
	}

	if user.Contact == "" {
		return &cerrors.EmptyValidationError{Field: "contact"}
	}

	if user.Address == "" {
		return &cerrors.EmptyValidationError{Field: "address"}
	}

	if user.Birth.IsZero() {
		return &cerrors.EmptyValidationError{Field: "birth"}
	}

	if err := validators.ValidateCPF(user.Cpf); err != nil {
		return err
	}

	return nil
}

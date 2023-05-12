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
	_, err := p.userServices.CreateUser(ctx, ofuscateData)
	if err != nil {
		return err
	}
	return nil
}

func (p *CreateUserService) Validate(user postgres.CreateUserParams) error {

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

func (p *CreateUserService) ObfuscateInformation(ctx context.Context, user postgres.CreateUserParams, fields []string) postgres.CreateUserParams {
	var obfuscateData postgres.CreateUserParams

	for _, s := range fields {
		switch s {
		case "Name":
			obfuscateData.Name = user.Name
			break
		case "Surname":
			obfuscateData.Surname = user.Surname
			break
		case "Contact":
			obfuscateData.Contact = user.Contact
			break
		case "Address":
			obfuscateData.Address = user.Address
			break
		case "Birth":
			obfuscateData.Birth = user.Birth
			break
		case "Cpf":
			obfuscateData.Cpf = user.Cpf
			break
		}
	}
	return obfuscateData
}

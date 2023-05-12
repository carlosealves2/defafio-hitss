package usecases

import (
	"context"
	"encoding/json"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
	"github.com/suportebeloj/desafio-hitss/src/protocols"
	"github.com/suportebeloj/desafio-hitss/src/utils/cerrors"
	"github.com/suportebeloj/desafio-hitss/src/utils/validators"
)

type CreateUserService struct {
	userServices   *postgres.Queries
	encryptService protocols.IEncrypterService
}

func NewProcessUserData(encryptService protocols.IEncrypterService) *CreateUserService {
	return &CreateUserService{
		encryptService: encryptService,
	}
}

func (p *CreateUserService) Create(ctx context.Context, user postgres.CreateUserParams) error {
	_, err := p.userServices.CreateUser(ctx, user)
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

// ObfuscateInformation this function encrypts sensitive data entered by the user through
// the fields argument (all fields must be written as they are defined in the model)
func (p *CreateUserService) ObfuscateInformation(ctx context.Context, user postgres.CreateUserParams, fields []string) (*postgres.CreateUserParams, error) {
	obfuscate := map[string]any{}

	useJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(useJson, &obfuscate); err != nil {
		return nil, err
	}
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		encoded, err := p.encryptService.Encrypt(obfuscate[field].(string))
		if err != nil {
			return nil, err
		}
		obfuscate[field] = encoded
	}

	obfuscateData := &postgres.CreateUserParams{}
	err = MapToStruct(obfuscate, obfuscateData)
	if err != nil {
		return nil, err
	}

	return obfuscateData, nil
}

func MapToStruct(m map[string]interface{}, s interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, s)
	if err != nil {
		return err
	}

	return nil
}

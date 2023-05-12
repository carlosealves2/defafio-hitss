package usecases

import (
	"github.com/stretchr/testify/assert"
	"github.com/suportebeloj/desafio-hitss/src/db/postgres"
	"github.com/suportebeloj/desafio-hitss/src/utils/cerrors"
	"testing"
	"time"
)

func TestProcessUserData_validate_ReturnFalse_WhenAnInvalidUserModelIsChecked(t *testing.T) {
	instance := NewProcessUserData()

	testData := []struct {
		expected error
		model    postgres.CreateUserParams
	}{
		{expected: cerrors.InvalidNameOrSurnameError{}, model: postgres.CreateUserParams{Surname: "TestSur"}},
		{expected: cerrors.InvalidNameOrSurnameError{}, model: postgres.CreateUserParams{Name: "TestName"}},
		{expected: cerrors.InvalidNameOrSurnameError{}, model: postgres.CreateUserParams{Name: "Test96"}},
		{expected: cerrors.InvalidNameOrSurnameError{}, model: postgres.CreateUserParams{Surname: "TestSur63"}},
		{expected: &cerrors.EmptyValidationError{Field: "contact"}, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: ""}},
		{expected: &cerrors.EmptyValidationError{Field: "address"}, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: "03493434", Address: ""}},
		{expected: &cerrors.EmptyValidationError{Field: "birth"}, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: "03434534593434", Address: "test address 123", Birth: time.Time{}, Cpf: ""}},
		{expected: cerrors.InvalidSizeCPFError{}, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: "03493434", Address: "test address 123", Birth: time.Now(), Cpf: ""}},
		{expected: cerrors.InvalidSizeCPFError{}, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: "03493434", Address: "test address 123", Birth: time.Now(), Cpf: "111.abc.222-22"}},
		{expected: cerrors.InvalidSizeCPFError{}, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: "03493434", Address: "test address 123", Birth: time.Now(), Cpf: "aged33.2344"}},
		{expected: nil, model: postgres.CreateUserParams{Name: "Test", Surname: "SurName", Contact: "03493434", Address: "test address 123", Birth: time.Now(), Cpf: "111.002.222-22"}},
	}

	for _, d := range testData {
		err := instance.validate(d.model)
		if err != nil {
			assert.EqualError(t, err, d.expected.Error())
		} else {
			assert.NoError(t, err)
		}

	}
}

package cerrors

import "fmt"

type EmptyValidationError struct {
	Field string
}

func (e *EmptyValidationError) Error() string {
	return fmt.Sprintf("%s cannot be empty", e.Field)
}

type InvalidNameOrSurnameError struct {
}

func (i InvalidNameOrSurnameError) Error() string {
	return "first or surname cannot be empty or contain numbers or special characters"
}

type InvalidSizeCPFError struct {
}

func (c InvalidSizeCPFError) Error() string {
	return "cpf cannot be greater or less than 11 numerical digits"
}

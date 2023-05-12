package validators

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateCPF(cpf string) error {
	pattern := `\d{11}`
	re, _ := regexp.Compile(`\D`)
	cleaned := re.ReplaceAllString(cpf, "")
	if len(cleaned) != 11 {
		return errors.New("cpf cannot be greater or less than 11 numerical digits")
	}

	re, _ = regexp.Compile(pattern)

	match := re.FindStringSubmatch(cleaned)
	if match == nil {
		return errors.New("entered cpf is not valid")
	}
	return nil
}

func InvalidChar(text string) bool {
	invalidNameChars := "1234567890'!\"#$%&\\'()*+,-./:;<=>?@[\\\\]^_`{|}~'"
	return strings.ContainsAny(text, invalidNameChars)
}

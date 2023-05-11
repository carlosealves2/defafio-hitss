package encrypter_test

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/suportebeloj/desafio-hitss/src/utils/encrypter"
	"os"
	"testing"
)

func init() {
	const keySize = 32

	key := make([]byte, keySize)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	os.Setenv("SECRET", string(key))
}

type EncryptSuiteTest struct {
	suite.Suite
}

func TestEncrypter_Encrypt_AString_WithSuccess_AndReturnABase64String(t *testing.T) {
	testPhrase := "plain text string"
	instance := encrypter.NewEncrypter()
	encoded, err := instance.Encrypt(testPhrase)
	assert.NoError(t, err)

	//	check if encoded string is base64
	_, err = base64.StdEncoding.DecodeString(encoded)
	assert.NoError(t, err)
}

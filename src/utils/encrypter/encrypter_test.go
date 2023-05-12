package encrypter_test

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"github.com/suportebeloj/desafio-hitss/src/utils/encrypter"
	"testing"
)

func TestEncrypter_Setup(t *testing.T) {
	instance := encrypter.NewEncrypter()
	assert.NotNil(t, instance)
}

func TestEncrypter_Encrypt_AString_WithSuccess(t *testing.T) {
	textToEncrypt := "encrypt text to test"

	instance := encrypter.NewEncrypter()
	assert.NotNil(t, instance)
	encoded, err := instance.Encrypt(textToEncrypt)
	assert.NoError(t, err)
	_, err = base64.StdEncoding.DecodeString(encoded)
	assert.NoError(t, err)
}

func TestEncrypter_Decrypt_AString_WithSuccess(t *testing.T) {
	textToEncrypt := "encrypt text to test"

	instance := encrypter.NewEncrypter()
	assert.NotNil(t, instance)
	encoded, err := instance.Encrypt(textToEncrypt)
	assert.NoError(t, err)
	_, err = base64.StdEncoding.DecodeString(encoded)
	assert.NoError(t, err)

	plainText, err := instance.Decrypt(encoded)
	assert.NoError(t, err)
	assert.Equal(t, plainText, textToEncrypt)

}

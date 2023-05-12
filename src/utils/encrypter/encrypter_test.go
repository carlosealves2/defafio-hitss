package tests_test

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"github.com/suportebeloj/desafio-hitss/src/utils/encrypter"
	"os"
	"testing"
)

func init() {
	key := generateFakeKey()
	os.Setenv("SECRET", string(key))
}

func generateFakeKey() (key []byte) {
	const keySize = 32

	key = make([]byte, keySize)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	return
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

func TestEncrypter_Decrypt_ABase64String_WithSuccess(t *testing.T) {
	testPhrase := "plain text string"
	instance := encrypter.NewEncrypter()
	encoded, err := instance.Encrypt(testPhrase)
	assert.NoError(t, err)

	decoded, err := instance.Decrypt(encoded)
	assert.NoError(t, err)
	assert.Equal(t, decoded, testPhrase)
}

func TestEncrypter_Decrypt_GivenAnError_WhenTryDecrypt_UsingAIvalidHash(t *testing.T) {
	testPhrase := "plain text string"
	instance := encrypter.NewEncrypter()
	encoded, err := instance.Encrypt(testPhrase)
	assert.NoError(t, err)

	encoded = string([]byte(encoded)[3:])
	_, err = instance.Decrypt(encoded)
	assert.Error(t, err)
}

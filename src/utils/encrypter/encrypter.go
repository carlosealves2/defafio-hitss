package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"
)

type Encrypter struct {
}

func NewEncrypter() *Encrypter {
	return &Encrypter{}
}

func (e Encrypter) Encrypt(plaintext string) (string, error) {
	key := []byte(os.Getenv("SECRET"))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	rawtext := []byte(plaintext)

	padded := make([]byte, len(rawtext)+(aes.BlockSize-len(rawtext)%aes.BlockSize))
	copy(padded, rawtext)

	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	final := append(iv, ciphertext...)

	encoded := base64.StdEncoding.EncodeToString(final)
	return encoded, nil
}

func (e Encrypter) Decrypt(encryptedText string) (string, error) {
	//TODO implement me
	panic("implement me")
}

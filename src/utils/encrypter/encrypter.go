package encrypter

import (
	"bytes"
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

	// Padding the text with padding to be a multiple of the block size
	padded := make([]byte, len(rawtext)+(aes.BlockSize-len(rawtext)%aes.BlockSize))
	copy(padded, rawtext)

	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	final := append(iv, ciphertext...)

	encoded := base64.StdEncoding.EncodeToString(final)
	return encoded, nil
}

func (e Encrypter) Decrypt(encryptedText string) (string, error) {
	key := []byte(os.Getenv("SECRET"))

	decoded, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := decoded[:aes.BlockSize]
	ciphertext := decoded[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove padding from text
	unpadded := bytes.TrimRight(plaintext, "\x00")
	return string(unpadded), nil
}

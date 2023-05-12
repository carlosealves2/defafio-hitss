package encrypter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	rsaKeySize = 2048
)

type keypair struct {
	priv *rsa.PrivateKey
	pub  *rsa.PublicKey
}

type Encrypter struct {
	kp            keypair
	signedMessage []byte
	hashed        [32]byte
	rng           io.Reader
	label         string
}

func NewEncrypter() *Encrypter {
	instance := &Encrypter{rng: rand.Reader}
	instance.setup()
	return instance
}

func (e *Encrypter) Loadkeys() {
	keyPath := os.Getenv("KEY_PATH")
	privKeyPath := fmt.Sprintf("%s/private_key.pem", keyPath)
	pubKeyPath := fmt.Sprintf("%s/public.pem", keyPath)
	if _, err := os.Stat(privKeyPath); os.IsNotExist(err) {
		e.GenerateKeyPair()
		e.SaveKeyPair()
		return
	}

	e.loadPrivateKey(privKeyPath)
	e.loadPublicKey(pubKeyPath)
}

func (e *Encrypter) loadPrivateKey(filename string) error {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return errors.New("failed to decode PEM block containing private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	e.kp.priv = key

	return nil
}

func (e *Encrypter) loadPublicKey(filename string) (*rsa.PublicKey, error) {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (e *Encrypter) GenerateKeyPair() error {
	var err error

	e.kp.priv, err = rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return err
	}

	e.kp.pub = &e.kp.priv.PublicKey
	return nil
}

func (e *Encrypter) SaveKeyPair() error {
	keyPath := os.Getenv("KEY_PATH")
	privKeyPath := fmt.Sprintf("%s/private_key.pem", keyPath)
	pubKeyPath := fmt.Sprintf("%s/public.pem", keyPath)

	keyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(e.kp.priv),
	})

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&e.kp.priv.PublicKey)
	if err != nil {
		return err
	}

	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	if err := os.WriteFile(privKeyPath, keyBytes, 0600); err != nil {
		return err
	}

	if err := os.WriteFile(pubKeyPath, pubKeyPEM, 0644); err != nil {
		return err
	}

	log.Println("private and public key generated and saved")

	return nil
}

func (e *Encrypter) Encrypt(plainText string) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), e.rng, e.kp.pub, []byte(plainText), []byte(e.label))
	encode := base64.StdEncoding.EncodeToString(ciphertext)
	if err != nil {
		return "", err
	}
	return encode, nil
}

func (e *Encrypter) Decrypt(encode string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptOAEP(sha256.New(), e.rng, e.kp.priv, decode, []byte(e.label))
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (e *Encrypter) setup() {
	e.Loadkeys()
	label := os.Getenv("SECRET_PHRASE")
	if label == "" {
		log.Fatalln("secret phrase not set on environment")
		return
	}

	e.label = label
}

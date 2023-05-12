package protocols

type IEncrypterService interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(encryptedText string) (string, error)
}

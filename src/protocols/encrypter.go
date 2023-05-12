package protocols

type IEncrypterService interface {
	Encrypt(plainText string) (string, error)
	Decrypt(encryptedText string) (string, error)
}

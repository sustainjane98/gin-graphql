package encryption

type EncryptionService struct {
}

func Bcrypt() BcryptService {
	return BcryptService{}
}

package services

import (
	"example/services/actions"
	"example/services/converter"
	"example/services/db"
	"example/services/encryption"
	"example/services/validators"
)

func DB() db.Services {
	return db.Services{}
}

func Validator() validators.ValidatorServices {
	return validators.ValidatorServices{}
}

func Encryption() encryption.EncryptionService {
	return encryption.EncryptionService{}
}

func DBActions() actions.ActionServices {
	return actions.ActionServices{}
}

func Converter() converter.ConverterServices {
	return converter.ConverterServices{}
}

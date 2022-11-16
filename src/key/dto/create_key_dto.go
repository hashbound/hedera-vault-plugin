package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateKeyDTO struct {
	ID        string `json:"id" validate:"required"`
	Algorithm string `json:"algorithm" validate:"required,oneof=ED25519 ECDSA"`
	Curve     string `json:"crve" validate:"oneof=secp256k1 ''"`
}

func (createKeyDTO *CreateKeyDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(createKeyDTO)
	if err != nil {
		return fmt.Errorf("validate create key parameters failed: %s", err)
	}

	return nil
}

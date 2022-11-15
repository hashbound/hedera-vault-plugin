package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ImportKeyDTO struct {
	ID         string `json:"id" validate:"required"`
	PrivateKey string `json:"privateKey" validate:"required"`
	Algorithm  string `json:"algorithm" validate:"required,oneof=ED25519 ECDSA"`
	Curve      string `json:"crve" validate:"required,oneof=secp256k1 ''"`
}

func (importKeyDTO *ImportKeyDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(importKeyDTO)
	if err != nil {
		return fmt.Errorf("validate create key parameters failed: %s", err)
	}

	return nil
}

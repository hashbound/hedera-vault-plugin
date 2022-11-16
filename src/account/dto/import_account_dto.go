package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ImportAccountDTO struct {
	ID        string `json:"id" validate:"required"`
	AccountID string `json:"accountId" validate:"required"`
	KeyID     string `json:"keyId" validate:"required"`
}

func (importAccountDTO *ImportAccountDTO) validate() error {
	err := validator.New().Struct(importAccountDTO)
	if err != nil {
		return fmt.Errorf("validate get account parameters failed")
	}

	return nil
}

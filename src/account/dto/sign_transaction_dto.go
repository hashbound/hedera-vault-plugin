package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SignTransactionDTO struct {
	ID          string `json:"accountId" validate:"required"`
	Transaction []byte `json:"transaction" validate:"required"`
}

func (signTransactionDTO *SignTransactionDTO) Validate() error {
	err := validator.New().Struct(signTransactionDTO)
	if err != nil {
		return fmt.Errorf("validate sign transaction parameters failed")
	}

	return nil
}

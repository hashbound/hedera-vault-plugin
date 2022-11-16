package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateAccountDTO struct {
	OriginID string `json:"originId" validate:"required"`
	ID       string `json:"id" validate:"required"`
}

func (createAccountDTO *CreateAccountDTO) validate() error {
	err := validator.New().Struct(createAccountDTO)
	if err != nil {
		return fmt.Errorf("validate create account parameters failed")
	}

	return nil
}

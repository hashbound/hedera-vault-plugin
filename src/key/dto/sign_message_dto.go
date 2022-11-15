package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SignMessageDTO struct {
	ID string `json:"id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func (signMessageDTO *SignMessageDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(signMessageDTO)
	if err != nil {
		return fmt.Errorf("validate create key parameters failed: %s", err)
	}

	return nil
}
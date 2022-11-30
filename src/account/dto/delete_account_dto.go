package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type DeleteAccountDTO struct {
	ID string `json:"id" validate:"required"`
}

func (deleteAccountDTO *DeleteAccountDTO) Validate() error {
	err := validator.New().Struct(deleteAccountDTO)
	if err != nil {
		return fmt.Errorf("validate delete account parameters failed")
	}

	return nil
}

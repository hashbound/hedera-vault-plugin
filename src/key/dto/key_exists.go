package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type KeyExistsDTO struct {
	ID string `json:"id" validate:"required"`
}

func (keyExistsDTO *KeyExistsDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(keyExistsDTO)
	if err != nil {
		return fmt.Errorf("validate retreive key parameters failed: %s", err)
	}

	return nil
}

package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type DeleteKeyDTO struct {
	ID string `json:"id" validate:"required"`
}

func (deleteKeyDTO *DeleteKeyDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(deleteKeyDTO)
	if err != nil {
		return fmt.Errorf("validate create key parameters failed: %s", err)
	}

	return nil
}

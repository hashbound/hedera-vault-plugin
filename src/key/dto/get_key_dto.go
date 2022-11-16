package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type GetKeyDTO struct {
	ID string `json:"id" validate:"required"`
}

func (getKeyDTO *GetKeyDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(getKeyDTO)
	if err != nil {
		return fmt.Errorf("validate create key parameters failed: %s", err)
	}

	return nil
}

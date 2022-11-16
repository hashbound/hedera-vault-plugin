package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type GetAccountDTO struct {
	ID string `json:"id" validate:"required"`
}

func (getAccountDTO *GetAccountDTO) Validate() error {
	err := validator.New().Struct(getAccountDTO)
	if err != nil {
		return fmt.Errorf("validate get account parameters failed")
	}

	return nil
}

package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type ImportAccountDTO struct {
	ID        string `json:"id" validate:"required"`
	AccountID string `json:"accountId" validate:"required"`
	KeyID     string `json:"keyId" validate:"required"`
}

func (importAccountDTO *ImportAccountDTO) Validate() error {
	err := validator.New().Struct(importAccountDTO)
	if err != nil {
		return fmt.Errorf("validate import account parameters failed")
	}

	// validate AccountID
	_, err = hedera.AccountIDFromString(importAccountDTO.AccountID)
	if err != nil {
		return fmt.Errorf("invalid account id: %s", err)
	}

	return nil
}

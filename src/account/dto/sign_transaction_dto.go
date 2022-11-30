package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type SignTransactionDTO struct {
	AccountID   string `json:"accountId" validate:"required"`
	Transaction string `json:"transaction" validate:"required"`
}

func (signTransactionDTO *SignTransactionDTO) Validate() error {
	err := validator.New().Struct(signTransactionDTO)
	if err != nil {
		return fmt.Errorf("validate sign transaction parameters failed")
	}

	// validate AccountID
	_, err = hedera.AccountIDFromString(signTransactionDTO.AccountID)
	if err != nil {
		return fmt.Errorf("invalid account id: %s", err)
	}

	return nil
}

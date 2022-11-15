package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type UnfreezeAccountDTO struct {
	accountID string `validate:"required"`
	kycKey    string `validate:"required"`
}

func (unfreezeAccountDTO *UnfreezeAccountDTO) validate() (*UnfreezeAccountParams, error) {
	validate := validator.New()
	err := validate.Struct(unfreezeAccountDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid unfreeze token parameters")
	}

	accountID, err := hedera.AccountIDFromString(unfreezeAccountDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	kycKey, err := hedera.PrivateKeyFromString(unfreezeAccountDTO.kycKey)
	if err != nil {
		return nil, fmt.Errorf("invalid KYC key: %s", err)
	}

	return &UnfreezeAccountParams{accountID, kycKey}, nil
}

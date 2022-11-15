package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type FreezeAccountDTO struct {
	accountID string `validate:"required"`
	kycKey    string `validate:"required"`
}

func (freezeAccountDTO *FreezeAccountDTO) validate() (*FreezeAccountParams, error) {
	validate := validator.New()
	err := validate.Struct(freezeAccountDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid freeze token parameters")
	}

	accountID, err := hedera.AccountIDFromString(freezeAccountDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	kycKey, err := hedera.PrivateKeyFromString(freezeAccountDTO.kycKey)
	if err != nil {
		return nil, fmt.Errorf("invalid KYC key: %s", err)
	}

	return &FreezeAccountParams{accountID, kycKey}, nil
}

package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type FreezeAccountDTO struct {
	accountID string
	kycKey    string
}

func (freezeAccountDTO *FreezeAccountDTO) validate() (*FreezeAccountParams, error) {
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
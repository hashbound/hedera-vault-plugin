package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type UnfreezeAccountDTO struct {
	accountID string
	kycKey    string
}

func (unfreezeAccountDTO *UnfreezeAccountDTO) validate() (*UnfreezeAccountParams, error) {
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
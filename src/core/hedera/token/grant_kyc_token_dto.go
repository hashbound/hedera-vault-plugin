package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type GrantKycDTO struct {
	accountID string
	kycKey    string
}

func (grantKycDTO *GrantKycDTO) validate() (*GrantKycParams, error) {
	accountID, err := hedera.AccountIDFromString(grantKycDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	kycKey, err := hedera.PrivateKeyFromString(grantKycDTO.kycKey)
	if err != nil {
		return nil, fmt.Errorf("invalid KYC key: %s", err)
	}

	return &GrantKycParams{accountID, kycKey}, nil
}
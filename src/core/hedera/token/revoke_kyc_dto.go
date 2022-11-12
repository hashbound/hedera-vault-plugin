package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type RevokeKycDTO struct {
	accountID string
	kycKey    string
}

func (revokeKycDTO *RevokeKycDTO) validate() (*RevokeKycParams, error) {
	accountID, err := hedera.AccountIDFromString(revokeKycDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	kycKey, err := hedera.PrivateKeyFromString(revokeKycDTO.kycKey)
	if err != nil {
		return nil, fmt.Errorf("invalid KYC key: %s", err)
	}

	return &RevokeKycParams{accountID, kycKey}, nil
}
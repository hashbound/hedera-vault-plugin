package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type RevokeKycDTO struct {
	accountID string `validate:"required"`
	kycKey    string `validate:"required"`
}

func (revokeKycDTO *RevokeKycDTO) validate() (*RevokeKycParams, error) {
	validate := validator.New()
	err := validate.Struct(revokeKycDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid revoke KYC parameters")
	}

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

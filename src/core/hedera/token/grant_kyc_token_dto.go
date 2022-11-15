package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type GrantKycDTO struct {
	accountID string `validate:"required"`
	kycKey    string `validate:"required"`
}

func (grantKycDTO *GrantKycDTO) validate() (*GrantKycParams, error) {
	validate := validator.New()
	err := validate.Struct(grantKycDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid grant KYC parameters")
	}

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

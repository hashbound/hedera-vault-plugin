package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type AssociateTokenDTO struct {
	accountID  string `validate:"required"`
	accountKey string `validate:"required"`
}

func (associateTokenDTO *AssociateTokenDTO) validate() (*AssociateWithTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(associateTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid associate token parameters")
	}

	accountID, err := hedera.AccountIDFromString(associateTokenDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	accountKey, err := hedera.PrivateKeyFromString(associateTokenDTO.accountKey)
	if err != nil {
		return nil, fmt.Errorf("invalid account key: %s", err)
	}

	return &AssociateWithTokenParams{accountID, accountKey}, nil
}

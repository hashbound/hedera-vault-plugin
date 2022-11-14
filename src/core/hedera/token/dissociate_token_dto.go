package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type DissociateTokenDTO struct {
	accountID  string	`validate:"required"`
	accountKey string	`validate:"required"`
}

func (dissociateTokenDTO *DissociateTokenDTO) validate() (*DissociateWithTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(dissociateTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid dissociate token parameters")
	}
	
	accountID, err := hedera.AccountIDFromString(dissociateTokenDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	accountKey, err := hedera.PrivateKeyFromString(dissociateTokenDTO.accountKey)
	if err != nil {
		return nil, fmt.Errorf("invalid account key: %s", err)
	}

	return &DissociateWithTokenParams{accountID, accountKey}, nil
}

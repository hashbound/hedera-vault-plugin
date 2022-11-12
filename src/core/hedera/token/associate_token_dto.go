package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type AssociateTokenDTO struct {
	accountID  string
	accountKey string
}

func (associateTokenDTO *AssociateTokenDTO) validate() (*AssociateWithTokenParams, error) {
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

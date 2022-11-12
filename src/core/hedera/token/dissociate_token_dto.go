package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type DissociateTokenDTO struct {
	accountID  string
	accountKey string
}

func (dissociateTokenDTO *DissociateTokenDTO) validate() (*DissociateWithTokenParams, error) {
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

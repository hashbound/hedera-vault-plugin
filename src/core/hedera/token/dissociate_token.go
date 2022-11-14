package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type DissociateWithTokenParams struct {
	accountID  hedera.AccountID
	accountKey hedera.PrivateKey
}

func (t *Token) DissociateWithToken(dissociateTokenDTO *DissociateTokenDTO) (*hedera.Status, error) {
	dissociateTokenParams, err := dissociateTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid dissociate token params: %s", err)
	}

	transaction, err := hedera.
		NewTokenDissociateTransaction().
		SetTokenIDs(t.TokenID).
		SetAccountID(dissociateTokenParams.accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(dissociateTokenParams.accountKey).
		Execute(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("execute transaction failed: %s", err)
	}

	receipt, err := response.GetReceipt(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("retreive transaction response failed: %s", err)
	}

	return &receipt.Status, nil
}

package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type AssociateWithTokenParams struct {
	accountID  hedera.AccountID
	accountKey hedera.PrivateKey
}

func (t *Token) AssociateWithToken(associateTokenDTO *AssociateTokenDTO) (*hedera.Status, error) {
	associateTokenParams, err := associateTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid associate token params: %s", err)
	}

	transaction, err := hedera.
		NewTokenAssociateTransaction().
		SetTokenIDs(t.TokenID).
		SetAccountID(associateTokenParams.accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(associateTokenParams.accountKey).
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

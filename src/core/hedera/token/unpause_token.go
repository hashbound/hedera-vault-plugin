package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type UnpauseTokenParams struct {
	pauseKey hedera.PrivateKey
}

func (t *Token) UnpauseToken(unpauseTokenDTO *UnpauseTokenDTO) (*hedera.Status, error) {
	unpauseTokenParams, err := unpauseTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid unpause token parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenUnpauseTransaction().
		SetTokenID(t.TokenID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(unpauseTokenParams.pauseKey).
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

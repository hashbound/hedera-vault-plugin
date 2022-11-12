package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type PauseTokenParams struct {
	pauseKey hedera.PrivateKey
}

func (t *Token) PauseToken(pauseTokenDTO *PauseTokenDTO) (*hedera.Status, error) {
	pauseTokenParams, err := pauseTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid pause token parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenPauseTransaction().
		SetTokenID(t.TokenID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(pauseTokenParams.pauseKey).
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

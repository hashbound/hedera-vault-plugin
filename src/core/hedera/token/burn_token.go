package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type BurnTokenParams struct {
	amount    uint64
	supplyKey *hedera.PrivateKey
}

func (t *Token) BurnToken(burnTokenDTO *BurnTokenDTO) (*hedera.Status, error) {
	burnTokenParams, err := burnTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid burn Token Parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenBurnTransaction().
		SetTokenID(t.TokenID).
		SetAmount(burnTokenParams.amount).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(*burnTokenParams.supplyKey).
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

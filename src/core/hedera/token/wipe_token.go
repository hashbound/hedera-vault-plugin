package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type WipeTokenParams struct {
	amount    uint64
	accountID hedera.AccountID
	wipeKey   hedera.PrivateKey
}

func (t *Token) WipeToken(wipeTokenDTO *WipeTokenDTO) (*hedera.Status, error) {
	wipeTokenParams, err := wipeTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid wipe token parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenWipeTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(wipeTokenParams.accountID).
		SetAmount(wipeTokenParams.amount).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(wipeTokenParams.wipeKey).
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

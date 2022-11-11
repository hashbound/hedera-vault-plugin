package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (t *Token) BurnToken(amount uint64, supplyKeyString string) (*hedera.Status, error) {
	supplyKey, err := hedera.PrivateKeyFromString(supplyKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenBurnTransaction().
		SetTokenID(t.TokenID).
		SetAmount(amount).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(supplyKey).
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

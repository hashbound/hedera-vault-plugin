package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (t *Token) PauseToken(puaseKeyString string) (*hedera.Status, error) {
	pauseKey, err := hedera.PrivateKeyFromString(puaseKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenPauseTransaction().
		SetTokenID(t.TokenID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(pauseKey).
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

package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (t *Token) WipeToken(accountIDString string, amount uint64, wipeKeyString string) (*hedera.Status, error) {
	accountID, err := hedera.AccountIDFromString(accountIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	wipeKey, err := hedera.PrivateKeyFromString(wipeKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenWipeTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(accountID).
		SetAmount(amount).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(wipeKey).
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

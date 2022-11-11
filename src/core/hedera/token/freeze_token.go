package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (t *Token) FreezeAccount(accountIDString, kycKeyString string) (*hedera.Status, error) {
	accountID, err := hedera.AccountIDFromString(accountIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	kycKey, err := hedera.PrivateKeyFromString(kycKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenFreezeTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(kycKey).
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

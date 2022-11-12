package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type FreezeAccountParams struct {
	accountID hedera.AccountID
	kycKey    hedera.PrivateKey
}

func (t *Token) FreezeAccount(freezeAccountDTO *FreezeAccountDTO) (*hedera.Status, error) {
	freezeAccountParams, err := freezeAccountDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid freeze token parameteers: %s", err)
	}

	transaction, err := hedera.
		NewTokenFreezeTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(freezeAccountParams.accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(freezeAccountParams.kycKey).
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

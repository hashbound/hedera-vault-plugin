package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type UnfreezeAccountParams struct {
	accountID hedera.AccountID
	kycKey    hedera.PrivateKey
}

func (t *Token) UnfreezeAccount(unfreezeTokenDTO *UnfreezeAccountDTO) (*hedera.Status, error) {
	unfreezeTokenParams, err := unfreezeTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid unfreeze token parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenUnfreezeTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(unfreezeTokenParams.accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(unfreezeTokenParams.kycKey).
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

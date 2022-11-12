package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type MintTokenParams struct {
	amount    uint64
	supplyKey *hedera.PrivateKey
}

func (t *Token) MintToken(mintTokenDTO *MintTokenDTO) (*hedera.Status, error) {
	mintTokenParams, err := mintTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid Mint Token Parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenMintTransaction().
		SetTokenID(t.TokenID).
		SetAmount(mintTokenParams.amount).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(*mintTokenParams.supplyKey).
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

package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type GrantKycParams struct {
	accountID hedera.AccountID
	kycKey    hedera.PrivateKey
}

func (t *Token) GrantKyc(grantKycDTO *GrantKycDTO) (*hedera.Status, error) {
	grantKycParams, err := grantKycDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid grant KYC parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenGrantKycTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(grantKycParams.accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(grantKycParams.kycKey).
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

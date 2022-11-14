package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type RevokeKycParams struct {
	accountID hedera.AccountID
	kycKey    hedera.PrivateKey
}

func (t *Token) RevokeKyc(revokeKycDTO *RevokeKycDTO) (*hedera.Status, error) {
	revokeKycParams, err := revokeKycDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid revoke KYC parameters: %s", err)
	}

	transaction, err := hedera.
		NewTokenRevokeKycTransaction().
		SetTokenID(t.TokenID).
		SetAccountID(revokeKycParams.accountID).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(revokeKycParams.kycKey).
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

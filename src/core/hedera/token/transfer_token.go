package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type TransferTokenParams struct {
	amount     uint64
	senderID   hedera.AccountID
	receiverID hedera.AccountID
	senderKey  hedera.PrivateKey
}

func (t *Token) TransferToken(trnasferTokenDTO *TransferTokenDTO) (*hedera.Status, error) {
	transferTokenParams, err := trnasferTokenDTO.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid transfer token parameters: %s", err)
	}

	transaction, err := hedera.NewTransferTransaction().
		AddTokenTransfer(t.TokenID, transferTokenParams.senderID, int64(-transferTokenParams.amount)).
		AddTokenTransfer(t.TokenID, transferTokenParams.receiverID, int64(transferTokenParams.amount)).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(transferTokenParams.senderKey).
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

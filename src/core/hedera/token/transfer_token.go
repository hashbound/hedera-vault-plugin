package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (t *Token) TransferToken(senderIDString, senderKeyString, recipientIDString string, amount uint64) (*hedera.Status, error) {
	senderID, err := hedera.AccountIDFromString(senderIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid senderID: %s", err)
	}
	senderKey, err := hedera.PrivateKeyFromString(senderKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalif senderKey: %s", err)
	}

	recipientID, err := hedera.AccountIDFromString(recipientIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid recipientID: %s", err)
	}

	transaction, err := hedera.NewTransferTransaction().
		AddTokenTransfer(t.TokenID, senderID, int64(amount)*-1).
		AddTokenTransfer(t.TokenID, recipientID, int64(amount)).
		FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(senderKey).
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

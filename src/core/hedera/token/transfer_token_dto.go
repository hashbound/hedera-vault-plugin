package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type TransferTokenDTO struct {
	amount     uint64
	senderID   string
	receiverID string
	senderKey  string
}

func (transferTokenDTO *TransferTokenDTO) validate() (*TransferTokenParams, error) {
	if transferTokenDTO.amount == 0 {
		return nil, fmt.Errorf("invalid transfer amount")
	}

	senderID, err := hedera.AccountIDFromString(transferTokenDTO.senderID)
	if err != nil {
		return nil, fmt.Errorf("invalid senderID: %s", err)
	}

	senderKey, err := hedera.PrivateKeyFromString(transferTokenDTO.senderKey)
	if err != nil {
		return nil, fmt.Errorf("invalif senderKey: %s", err)
	}

	receiverID, err := hedera.AccountIDFromString(transferTokenDTO.receiverID)
	if err != nil {
		return nil, fmt.Errorf("invalid receiverID: %s", err)
	}

	return &TransferTokenParams{
		amount: transferTokenDTO.amount,
		senderID: senderID,
		senderKey: senderKey,
		receiverID: receiverID,
	}, nil
}
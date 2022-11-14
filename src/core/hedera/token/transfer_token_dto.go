package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type TransferTokenDTO struct {
	amount     uint64	`validate:"required,gt=0"`
	senderID   string	`validate:"required"`
	receiverID string	`validate:"required"`
	senderKey  string	`validate:"required"`
}

func (transferTokenDTO *TransferTokenDTO) validate() (*TransferTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(transferTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid transfer token parameters")
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
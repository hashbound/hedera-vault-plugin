package gateway

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Operator struct {
	operatorID  hedera.AccountID
	operatorKey hedera.PrivateKey
}

func NewOperator(accountID, privateKey string) (*Operator, error) {
	operatorID, err := hedera.AccountIDFromString(accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid operatorID: %s", err)
	}

	operatorKey, err := hedera.PrivateKeyFromString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid operator key: %s", err)
	}

	return &Operator{operatorID, operatorKey}, nil
}

package gateway

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Operator struct {
	operatorID  string
	operatorKey string
}

func NewOperator(operatorID, operatorKey string) *Operator {
	return &Operator{operatorID, operatorKey}
}

func (gw *Gateway) WithOperator(operatorID, operatorKey string) *Gateway {
	gw.operator.operatorID = operatorID
	gw.operator.operatorKey = operatorKey

	return gw
}

func (gw *Gateway) SetOperator() (*Gateway, error) {
	operatorID, err := hedera.AccountIDFromString(gw.operator.operatorID)
	if err != nil {
		return nil, fmt.Errorf("invalid operatorID: %s", err)
	}

	operatorKey, err := hedera.PrivateKeyFromString(gw.operator.operatorKey)
	if err != nil {
		return nil, fmt.Errorf("invalid operator key: %s", err)
	}

	gw.client = gw.client.SetOperator(operatorID, operatorKey)

	return gw, nil
}

func (gw *Gateway) SetWithOperator(operatorID, operatorKey string) (*Gateway, error) {
	return gw.WithOperator(operatorID, operatorKey).SetOperator()
}

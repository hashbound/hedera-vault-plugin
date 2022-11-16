package gateway

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Gateway struct {
	client   *hedera.Client
	operator *Operator
}

func New() *Gateway {
	return &Gateway{}
}

func (gw *Gateway) GetClient() *hedera.Client {
	return gw.client
}

func (gw *Gateway) WithClient(client *hedera.Client) *Gateway {
	gw.client = client
	return gw
}

func (gw *Gateway) WithOperator(operator *Operator) *Gateway {
	gw.operator = operator
	return gw
}

func (gw *Gateway) Set() (*Gateway, error) {
	gw.client = gw.client.SetOperator(gw.operator.operatorID, gw.operator.operatorKey)
	return gw, nil
}

func (gw *Gateway) SetWithOperator(operator *Operator) (*Gateway, error) {
	return gw.WithOperator(operator).Set()
}

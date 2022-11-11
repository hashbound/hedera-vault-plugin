package gateway

import "github.com/hashgraph/hedera-sdk-go/v2"

type Gateway struct {
	client   *hedera.Client
	operator *Operator
}

func New() *Gateway {
	return &Gateway{}
}

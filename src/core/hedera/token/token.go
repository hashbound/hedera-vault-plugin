package token

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/hedera/gateway"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Token struct {
	gateway *gateway.Gateway
	TokenID hedera.TokenID
}

func New(gateway *gateway.Gateway) *Token {
	return &Token{gateway: gateway}
}

func (t *Token) WithTokenID(tokenID hedera.TokenID) *Token {
	t.TokenID = tokenID
	return t
}

func (t *Token) WithTokenIDString(tokenID string) (*Token, error) {
	tID, err := hedera.TokenIDFromString(tokenID)
	if err != nil {
		return nil, fmt.Errorf("invalid tokenID: %s", err)
	}
	t.TokenID = tID

	return t, nil
}

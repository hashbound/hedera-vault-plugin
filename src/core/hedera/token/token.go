package token

import (
	"github.com/hashbound/hedera-vault-plugin/src/core/hedera/gateway"
)

type Token struct {
	gateway *gateway.Gateway
}

func New(gateway *gateway.Gateway) *Token {
	return &Token{gateway}
}

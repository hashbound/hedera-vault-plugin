package account

import (
	gw "github.com/hashbound/hedera-vault-plugin/src/core/hedera/gateway"
)

type Account struct {
	gateway *gw.Gateway
}

func New(gateway *gw.Gateway) *Account {
	return &Account{gateway}
}

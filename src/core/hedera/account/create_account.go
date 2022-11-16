package account

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (a *Account) CreateAccount(priv hedera.PrivateKey) (*hedera.AccountID, error) {
	tr, err := hedera.
		NewAccountCreateTransaction().
		SetKey(priv).
		Execute(a.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("create NewAccountTransaction failed: %s", err)
	}

	receipt, err := tr.GetReceipt(a.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("get receipt failed: %s", err)
	}

	return receipt.AccountID, nil
}

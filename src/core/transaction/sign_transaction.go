package transaction

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

func SignTransaction(rawTransaction []byte, privateKey *key.PrivateKey) ([]byte, error) {
	tnx, err := hedera.TransactionFromBytes(rawTransaction)
	if err != nil {
		return nil, fmt.Errorf("retreive transaction failed: %s", err)
	}

	keys, err := key.FromPrivateKey(*privateKey)
	if err != nil {
		return nil, fmt.Errorf("retreive private key failed: %s", err)
	}

	signedTransaction, err := hedera.TransactionSign(tnx, keys.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("sign transaction failed: %s", err)
	}

	signedTransactionByte, err := hedera.TransactionToBytes(signedTransaction)
	if err != nil {
		return nil, fmt.Errorf("convert transaction to bytes failed: %s", err)
	}

	return signedTransactionByte, nil
}

package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashbound/hedera-vault-plugin/src/core/transaction"
	key_dto "github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func (svc *AccountService) SignTransaction(signTransactionDTO *dto.SignTransactionDTO) ([]byte, error) {
	err := signTransactionDTO.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate sign transaction parameters failed: %s", err)
	}

	account, err := svc.storage.Read(signTransactionDTO.ID)
	if err != nil {
		return nil, fmt.Errorf("retreive account from storage failed: %s", err)
	}

	keypair, err := svc.k_svc.GetKey(&key_dto.GetKeyDTO{ID: account.KeyID})
	if err != nil {
		return nil, fmt.Errorf("retreive key from storage failed: %s", err)
	}

	privateKey, err := key.FromPrivateKey(key.PrivateKey{
		Key:       keypair.PrivateKey,
		Algorithm: key.AlgorithmFromString(keypair.Algorithm),
		Curve:     key.CurveFromString(keypair.Curve),
	})
	if err != nil {
		return nil, fmt.Errorf("retreive private key failed: %s", err)
	}

	signedTransaction, err := transaction.SignTransaction([]byte(signTransactionDTO.Transaction), privateKey.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("sign transaction failed: %s", err)
	}

	return signedTransaction, nil
}

package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashbound/hedera-vault-plugin/src/account/entity"
	hedera_account "github.com/hashbound/hedera-vault-plugin/src/core/hedera/account"
	"github.com/hashbound/hedera-vault-plugin/src/core/hedera/gateway"
	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	key_dto "github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func (svc *AccountService) CreateAccount(createAccountDTO *dto.CreateAccountDTO) (*entity.Account, error) {
	err := createAccountDTO.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate create account failed: %s", err)
	}

	originAccount, err := svc.storage.Read(createAccountDTO.OriginID)
	if err != nil {
		return nil, err
	}
	accountKey, err := svc.k_svc.GetKey(&key_dto.GetKeyDTO{ID: originAccount.KeyID})
	if err != nil {
		return nil, fmt.Errorf("retreive key failed: %s", err)
	}

	operator, err := gateway.NewOperator(originAccount.AccountID, accountKey.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("prepare gateway operator failed: %s", err)
	}
	gw, err := gateway.New().WithOperator(operator).ClientFromNetworkName("")
	if err != nil {
		return nil, fmt.Errorf("initialize gateway failed: %s", err)
	}

	ha := hedera_account.New(gw)
	new_account_id, err := ha.CreateAccount(key.PrivateKey{
		Key:       accountKey.PrivateKey,
		Algorithm: key.AlgorithmFromString(accountKey.Algorithm),
		Curve:     key.CurveFromString(accountKey.Curve),
	})
	if err != nil {
		return nil, fmt.Errorf("create new account failed: %s", err)
	}

	new_account := entity.New(createAccountDTO.OriginID, new_account_id.String(), originAccount.KeyID)
	if err := svc.storage.Write(createAccountDTO.OriginID, new_account); err != nil {
		return nil, fmt.Errorf("write account failed: %s", err)
	}

	return new_account, nil
}

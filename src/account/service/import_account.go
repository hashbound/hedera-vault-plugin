package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashbound/hedera-vault-plugin/src/account/entity"
	key_dto "github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func (svc *AccountService) ImportAccount(importAccountDTO *dto.ImportAccountDTO) (*entity.Account, error) {
	err := importAccountDTO.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate import account parameters failed: %s", err)
	}

	// verify if key exists with KeyID
	keyExists, err := svc.k_svc.KeyExists(&key_dto.KeyExistsDTO{
		ID: importAccountDTO.KeyID,
	})
	if err != nil {
		return nil, fmt.Errorf("retrive account key failed: %s", err)
	}
	if !keyExists {
		return nil, fmt.Errorf("key not found: %s", err)
	}

	account := entity.New(
		importAccountDTO.ID,
		importAccountDTO.AccountID,
		importAccountDTO.KeyID,
	)
	if err := svc.storage.
		Write(importAccountDTO.ID, account); err != nil {
		return nil, fmt.Errorf("import account failed: %s", err)
	}

	return account, nil
}

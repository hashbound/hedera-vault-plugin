package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashbound/hedera-vault-plugin/src/account/entity"
)

func (svc *AccountService) GetAccount(getAccountDTO *dto.GetAccountDTO) (*entity.Account, error) {
	err := getAccountDTO.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate get account failed: %s", err)
	}

	entity, err := svc.storage.Read(getAccountDTO.ID)
	if err != nil {
		return nil, fmt.Errorf("retreive account from storage failed: %s", err)
	}

	return entity, nil
}

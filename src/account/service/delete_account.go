package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
)

func (svc *AccountService) DeleteAccount(deleteAccountDTO *dto.DeleteAccountDTO) error {
	err := deleteAccountDTO.Validate()
	if err != nil {
		return fmt.Errorf("validate delete account parameters failed: %s", err)
	}

	entity, err := svc.storage.Read(deleteAccountDTO.ID)
	if err != nil {
		return fmt.Errorf("retreive account before deletion failed: %s", err)
	}
	if entity == nil {
		return fmt.Errorf("account does not exist")
	}

	return svc.storage.Delete(deleteAccountDTO.ID)
}

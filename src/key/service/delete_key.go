package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func (svc *KeyService) Delete(deleteKeyDTO *dto.DeleteKeyDTO) error {
	err := deleteKeyDTO.Validate()
	if err != nil {
		return fmt.Errorf("validate delete key parameters failed: %s", err)
	}

	entity, err := svc.storage.Read(deleteKeyDTO.ID)
	if err != nil {
		return fmt.Errorf("retreive key before deletion failed: %s", err)
	}
	if entity == nil {
		return fmt.Errorf("key does not exist")
	}

	return svc.storage.Delete(deleteKeyDTO.ID)
}

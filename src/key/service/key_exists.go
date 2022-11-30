package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func (k_svc *KeyService) KeyExists(keyExistsDTO *dto.KeyExistsDTO) (bool, error) {
	err := keyExistsDTO.Validate()
	if err != nil {
		return false, fmt.Errorf("validate key exists failed: %s", err)
	}

	entity, err := k_svc.storage.Read(keyExistsDTO.ID)
	if err != nil {
		return false, fmt.Errorf("retreive key from storage failed: %s", err)
	}

	if entity == nil {
		return false, nil
	}

	return true, nil
}

package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
	"github.com/hashbound/hedera-vault-plugin/src/key/entity"
)

func (k_svc *KeyService) GetKey(getKeyDTO *dto.GetKeyDTO) (*entity.Key, error) {
	err := getKeyDTO.Validate(); if err != nil {
		return nil, fmt.Errorf("validate get key failed: %s", err)
	}

	entity, err := k_svc.storage.Read(getKeyDTO.ID)
	if err != nil {
		return nil, fmt.Errorf("retreive key from storage failed: %s", err)
	}

	return entity, nil
}

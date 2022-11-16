package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
	entity "github.com/hashbound/hedera-vault-plugin/src/key/entity"
)

func (svc *KeyService) Create(createKeyDTO *dto.CreateKeyDTO) (*entity.Key, error) {
	err := createKeyDTO.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate create key failed: %s", err)
	}

	algo := key.AlgorithmFromString(createKeyDTO.Algorithm)
	curve := key.CurveFromString(createKeyDTO.Curve)
	keypair, err := key.GenerateKeyPair(algo, curve)
	if err != nil {
		return nil, fmt.Errorf("generate key pair failed: %s", err)
	}

	entity := entity.FromKeyPair(createKeyDTO.ID, keypair)
	if err := svc.storage.Write(createKeyDTO.ID, entity); err != nil {
		return nil, fmt.Errorf("write key to storage failed: %s", err)
	}

	return entity, nil
}

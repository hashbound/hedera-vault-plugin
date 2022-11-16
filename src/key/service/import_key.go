package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
	"github.com/hashbound/hedera-vault-plugin/src/key/entity"
)

func (svc *KeyService) ImportKey(importKey *dto.ImportKeyDTO) (*entity.Key, error) {
	err := importKey.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate create key failed: %s", err)
	}

	privateKey := key.PrivateKey{
		Key:       importKey.PrivateKey,
		Algorithm: key.AlgorithmFromString(importKey.Algorithm),
	}
	if importKey.Algorithm == key.Algorithm(key.ECDSA).String() {
		privateKey.Curve = key.CurveFromString(importKey.Curve)
	}

	keypair, err := key.FromPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("retreive key pair failed: %s", err)
	}

	entity := entity.FromKeyPair(importKey.ID, keypair)
	if err := svc.storage.Write(importKey.ID, entity); err != nil {
		return nil, fmt.Errorf("write key to storage failed: %s", err)
	}

	return entity, err
}

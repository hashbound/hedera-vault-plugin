package service

import (
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func (svc *KeyService) Sign(signMessageDTO *dto.SignMessageDTO) ([]byte, error) {
	err := signMessageDTO.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate sign message parameters failed: %s", err)
	}
	entity, err := svc.storage.Read(signMessageDTO.ID)
	if err != nil {
		return nil, fmt.Errorf("read key by ID failed: %s", err)
	}

	privateKey := key.PrivateKey{
		Key: entity.PrivateKey,
		Algorithm: key.AlgorithmFromString(entity.Algorithm),
	}
	if entity.Algorithm == key.Algorithm(key.ECDSA).String() {
		privateKey.Curve =  key.CurveFromString(entity.Curve)
	}

	signature, err := key.Sign(privateKey, []byte(signMessageDTO.Message))
	if err != nil {
		return nil, fmt.Errorf("signing message failed: %s", err)
	}

	return signature, nil
}

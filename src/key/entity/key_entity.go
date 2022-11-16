package entity

import (
	"encoding/json"

	"github.com/hashbound/hedera-vault-plugin/src/core/key"
)

type Key struct {
	ID         string `json:"id"`
	Algorithm  string `json:"algorithm"`
	Curve      string `json:"curve"`
	PrivateKey string `json:"privateKey"`
	Publickey  string `json:"publicKey"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func FromKeyPair(id string, kp *key.KeyPair) *Key {
	return &Key{
		ID:         id,
		Algorithm:  kp.Algorithm.String(),
		Curve:      kp.Curve.String(),
		PrivateKey: kp.PrivateKey.String(),
		Publickey:  kp.PublicKey.String(),
	}
}

func FromBytes(buf []byte) (*Key, error) {
	var key Key

	if err := json.Unmarshal(buf, &key); err != nil {
		return &Key{}, err
	}
	return &key, nil
}

func (k *Key) ToBytes() ([]byte, error) {
	val, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	return val, nil
}

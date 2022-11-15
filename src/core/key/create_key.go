package key

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type PrivateKey struct {
	Key       string
	Algorithm Algorithm
	Curve     Curve
}

type PublicKey struct {
	Key       string
	Algorithm Algorithm
	Curve     Curve
}

type KeyPair struct {
	PublicKey  hedera.PublicKey
	PrivateKey hedera.PrivateKey
	Algorithm  Algorithm
	Curve      Curve
}

func NewKeyPair(pub hedera.PublicKey, priv hedera.PrivateKey, algo Algorithm, curve Curve) *KeyPair {
	return &KeyPair{
		PublicKey:  pub,
		PrivateKey: priv,
		Algorithm:  algo,
		Curve:      curve,
	}
}

func GenerateKeyPair(algo Algorithm, curve Curve) (*KeyPair, error) {
	var priv hedera.PrivateKey
	var err error

	if algo == ED25519 {
		priv, err = hedera.PrivateKeyGenerateEd25519()
	} else if algo == ECDSA && curve == secp256k1 {
		priv, err = hedera.PrivateKeyGenerateEcdsa()
	} else {
		return &KeyPair{}, fmt.Errorf("invalid algorithm or curve")
	}

	if err != nil {
		return &KeyPair{}, err
	}

	pub := priv.PublicKey()
	return NewKeyPair(pub, priv, algo, curve), nil
}

func FromPrivateKey(privateKey PrivateKey) (*KeyPair, error) {
	var priv hedera.PrivateKey
	var err error

	if privateKey.Algorithm == ED25519 {
		priv, err = hedera.PrivateKeyFromStringEd25519(privateKey.Key)
	} else if privateKey.Algorithm == ECDSA && privateKey.Curve == secp256k1 {
		priv, err = hedera.PrivateKeyFromStringECSDA(privateKey.Key)
	} else {
		return nil, fmt.Errorf("invalid algorithm or curve")
	}

	if err != nil {
		return nil, fmt.Errorf("invalid private key: %s", err)
	}

	pub := priv.PublicKey()

	return NewKeyPair(pub, priv, privateKey.Algorithm, privateKey.Curve), err
}

func FromPublicKey(publicKey PublicKey) (*KeyPair, error) {
	var pub hedera.PublicKey
	var err error

	if publicKey.Algorithm == ED25519 {
		pub, err = hedera.PublicKeyFromStringEd25519(publicKey.Key)
	} else if publicKey.Algorithm == ECDSA && publicKey.Curve == secp256k1 {
		pub, err = hedera.PublicKeyFromStringECDSA(publicKey.Key)
	} else {
		return nil, fmt.Errorf("invalid algorithm or curve")
	}

	if err != nil {
		return nil, fmt.Errorf("invalid public key: %s", err)
	}

	return NewKeyPair(pub, hedera.PrivateKey{}, publicKey.Algorithm, publicKey.Curve), err
}

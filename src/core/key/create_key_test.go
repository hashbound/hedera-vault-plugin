package key

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPairED25519(t *testing.T) {
	key, err := GenerateKeyPair(ED25519, None)
	if err != nil {
		t.Fatalf("Unable to generate key pair %s", err)
	}

	fmt.Printf("Algorithm %s\n", key.Algorithm)
	fmt.Printf("Private Key %s\n", key.PrivateKey)
	fmt.Printf("Public Key %s\n", key.PublicKey)
}

func TestGenerateKeyPairECDSA(t *testing.T) {
	key, err := GenerateKeyPair(ECDSA, secp256k1)
	if err != nil {
		t.Fatalf("Unable to generate key pair %s", err)
	}

	fmt.Printf("Algorithm %s\n", key.Algorithm)
	fmt.Printf("Private Key %s\n", key.PrivateKey)
	fmt.Printf("Public Key %s\n", key.PublicKey)
}

func TestFromPrivateKeyED25519(t *testing.T) {
	privateKey := "302e020100300506032b657004220420d937fc715aed14a0433be15c1d7d66cea413837274e32787e717e220d21965fa"
	publicKey := "302a300506032b657003210027a0ab3727c10325d2ef39ca9f131814b0be4cf2037071b53eb7fd8afd164b3f"
	algo := ED25519

	key, err := FromPrivateKey(PrivateKey{
		Key:       privateKey,
		Algorithm: Algorithm(algo),
		Curve:     None,
	})
	if err != nil {
		t.Fatalf("Unable to derive key pair from string private key: %s\n", err)
	}

	fmt.Printf("Public Key: %s", key.PublicKey.String())

	assert.Equal(t, publicKey, key.PublicKey.String(), fmt.Sprintf("expected %s\nreceived: %s", publicKey, key.PublicKey.String()))
}

func TestFromPrivateKeyECDSA(t *testing.T) {
	privateKey := "3030020100300706052b8104000a04220420c16da8af48c66697f85c8cd05e7886449fd744555b9f688a9e2e6e66cbf8fce3"
	publicKey := "302f300706052b8104000a032400042102966386144b51b6b22b10044a9e878c1556b565325e0f0dae5492b598d18eba5e"
	algo := ECDSA
	curve := secp256k1

	key, err := FromPrivateKey(PrivateKey{
		Key:       privateKey,
		Algorithm: Algorithm(algo),
		Curve:     Curve(curve),
	})
	if err != nil {
		t.Fatalf("Unable to derive key pair from string private key: %s\n", err)
	}

	fmt.Printf("Public Key: %s", key.PublicKey.String())

	assert.Equal(t, publicKey, key.PublicKey.String(), fmt.Sprintf("expected %s\nreceived: %s", publicKey, key.PublicKey.String()))
}

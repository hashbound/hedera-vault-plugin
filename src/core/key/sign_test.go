package key

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSignED25519(t *testing.T) {
	privateKey := "302e020100300506032b657004220420d937fc715aed14a0433be15c1d7d66cea413837274e32787e717e220d21965fa"
	algo := ED25519
	digest := "test message"

	sig, err := Sign(PrivateKey{
		Key:       privateKey,
		Algorithm: Algorithm(algo),
	}, []byte(digest))
	if err != nil {
		t.Fatalf("unable to sign digest: %s", err)
	}

	fmt.Printf("Signature: %v", hex.EncodeToString(sig))
}

func TestSignECDSA(t *testing.T) {
	privateKey := "3030020100300706052b8104000a04220420c16da8af48c66697f85c8cd05e7886449fd744555b9f688a9e2e6e66cbf8fce3"
	algo := ECDSA
	curve := secp256k1
	digest := "test message"

	sig, err := Sign(PrivateKey{
		Key:       privateKey,
		Algorithm: Algorithm(algo),
		Curve:     Curve(curve),
	}, []byte(digest))
	if err != nil {
		t.Fatalf("unable to sign digest: %s", err)
	}

	fmt.Printf("Signature: %v", hex.EncodeToString(sig))
}

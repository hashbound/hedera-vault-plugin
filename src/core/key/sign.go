package key

import (
	"fmt"
)

func Sign(privateKey string, algo Algorithm, curve Curve, message []byte) ([]byte, error) {
	key, err := FromPrivateKey(PrivateKey{
		Key:       privateKey,
		Algorithm: algo,
		Curve:     curve,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid key or parameters: %s", err)
	}

	signature := key.PrivateKey.Sign([]byte(message))
	return signature, nil
}

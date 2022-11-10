package key

import (
	"fmt"
)

func Sign(privateKey, algo, curve string, message []byte) ([]byte, error) {
	key, err := FromPrivateKey(privateKey, algo, curve)
	if err != nil {
		return nil, fmt.Errorf("invalid key or parameters: %s", err)
	}

	signature := key.PrivateKey.Sign([]byte(message))
	return signature, nil
}

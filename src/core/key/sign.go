package key

import (
	"fmt"
)

func Sign(privateKey PrivateKey, message []byte) ([]byte, error) {
	priv, err := FromPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid key or parameters: %s", err)
	}

	signature := priv.PrivateKey.Sign([]byte(message))
	return signature, nil
}

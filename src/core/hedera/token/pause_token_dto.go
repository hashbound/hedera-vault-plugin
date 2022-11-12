package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type PauseTokenDTO struct {
	pauseKey string
}

func (pauseTokenDTO *PauseTokenDTO) validate() (*PauseTokenParams, error) {
	pauseKey, err := hedera.PrivateKeyFromString(pauseTokenDTO.pauseKey)
	if err != nil {
		return nil, fmt.Errorf("invalid pause key: %s", err)
	}

	return &PauseTokenParams{pauseKey}, nil
}
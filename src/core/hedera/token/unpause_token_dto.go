package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type UnpauseTokenDTO struct {
	pauseKey string
}

func (unpauseTokenDTO *UnpauseTokenDTO) validate() (*UnpauseTokenParams, error) {
	pauseKey, err := hedera.PrivateKeyFromString(unpauseTokenDTO.pauseKey)
	if err != nil {
		return nil, fmt.Errorf("invalid pause key: %s", err)
	}

	return &UnpauseTokenParams{pauseKey}, nil
}
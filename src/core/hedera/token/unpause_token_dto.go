package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type UnpauseTokenDTO struct {
	pauseKey string `validate:"required"`
}

func (unpauseTokenDTO *UnpauseTokenDTO) validate() (*UnpauseTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(unpauseTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid unpause token parameters")
	}

	pauseKey, err := hedera.PrivateKeyFromString(unpauseTokenDTO.pauseKey)
	if err != nil {
		return nil, fmt.Errorf("invalid pause key: %s", err)
	}

	return &UnpauseTokenParams{pauseKey}, nil
}

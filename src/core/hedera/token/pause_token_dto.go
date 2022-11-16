package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type PauseTokenDTO struct {
	pauseKey string `validate:"required"`
}

func (pauseTokenDTO *PauseTokenDTO) validate() (*PauseTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(pauseTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid pause token parameters")
	}

	pauseKey, err := hedera.PrivateKeyFromString(pauseTokenDTO.pauseKey)
	if err != nil {
		return nil, fmt.Errorf("invalid pause key: %s", err)
	}

	return &PauseTokenParams{pauseKey}, nil
}

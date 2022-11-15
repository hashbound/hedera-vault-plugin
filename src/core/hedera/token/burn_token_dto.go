package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type BurnTokenDTO struct {
	amount    uint64 `validate:"required,gt=0"`
	supplyKey string `validate:"required"`
}

func (burnTokenDTO *BurnTokenDTO) validate() (*BurnTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(burnTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid burn token parameters")
	}

	supplyKey, err := hedera.PrivateKeyFromString(burnTokenDTO.supplyKey)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	return &BurnTokenParams{
		amount:    burnTokenDTO.amount,
		supplyKey: &supplyKey,
	}, nil
}

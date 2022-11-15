package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type MintTokenDTO struct {
	amount    uint64 `validate:"required,gt=0"`
	supplyKey string `validate:"required"`
}

func (mintTokenDTO *MintTokenDTO) validate() (*MintTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(mintTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid mint token parameters")
	}

	supplyKey, err := hedera.PrivateKeyFromString(mintTokenDTO.supplyKey)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	return &MintTokenParams{
		amount:    mintTokenDTO.amount,
		supplyKey: &supplyKey,
	}, nil
}

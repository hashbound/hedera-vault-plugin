package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type BurnTokenDTO struct {
	amount    uint64
	supplyKey string
}

func (burnTokenDTO *BurnTokenDTO) validate() (*BurnTokenParams, error) {
	if burnTokenDTO.amount == 0 {
		return nil, fmt.Errorf("invalid amount value")
	}

	supplyKey, err := hedera.PrivateKeyFromString(burnTokenDTO.supplyKey)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	return &BurnTokenParams{
		amount: burnTokenDTO.amount,
		supplyKey: &supplyKey,
	}, nil
}
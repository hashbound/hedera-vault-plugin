package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type MintTokenDTO struct {
	amount    uint64
	supplyKey string
}

func (mintTokenDTO *MintTokenDTO) validate() (*MintTokenParams, error) {
	if mintTokenDTO.amount == 0 {
		return nil, fmt.Errorf("invalid amount value")
	}

	supplyKey, err := hedera.PrivateKeyFromString(mintTokenDTO.supplyKey)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	return &MintTokenParams{
		amount: mintTokenDTO.amount,
		supplyKey: &supplyKey,
	}, nil
}

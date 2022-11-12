package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type WipeTokenDTO struct {
	amount    uint64
	accountID string
	wipeKey   string
}

func (wipeTokenDTO *WipeTokenDTO) validate() (*WipeTokenParams, error) {
	if wipeTokenDTO.amount == 0 {
		return nil, fmt.Errorf("invalid wipe amount value")
	}

	accountID, err := hedera.AccountIDFromString(wipeTokenDTO.accountID)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	wipeKey, err := hedera.PrivateKeyFromString(wipeTokenDTO.wipeKey)
	if err != nil {
		return nil, fmt.Errorf("invalid wipe key: %s", err)
	}

	return &WipeTokenParams{
		amount: wipeTokenDTO.amount,
		accountID: accountID,
		wipeKey: wipeKey,
	}, nil
}
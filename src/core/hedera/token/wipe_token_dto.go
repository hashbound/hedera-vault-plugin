package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type WipeTokenDTO struct {
	amount    uint64	`validate:"required,gt=0"`
	accountID string	`validate:"required"`
	wipeKey   string	`validate:"required"`
}

func (wipeTokenDTO *WipeTokenDTO) validate() (*WipeTokenParams, error) {
	validate := validator.New()
	err := validate.Struct(wipeTokenDTO)
	if err != nil {
		return nil, fmt.Errorf("invalid wipe token parameters")
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
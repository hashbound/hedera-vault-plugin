package token

import (
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type CreateFungibleToken struct {
	Type               hedera.TokenType
	Name               string
	Symbol             string
	Decimals           uint
	InitSupply         uint64
	TreasuryAccountID  hedera.AccountID
	TreasuryPrivateKey hedera.PrivateKey
	AdminPublicKey     hedera.PublicKey
	AdminPrivateKey    hedera.PrivateKey
	KycKey             hedera.PublicKey
	FreezeKey          hedera.PublicKey
	WipeKey            hedera.PublicKey
	SupplyKey          hedera.PublicKey
	FeeScheduleKey     hedera.PublicKey
	PauseKey           hedera.PublicKey
	// CustomFees         []hedera.Fee
	MaxSupply        int64
	SupplyType       hedera.TokenSupplyType
	FreezeDefault    bool
	ExpirationTime   time.Time
	AutoRenewAccount hedera.AccountID
	Memo             string
}

func (t *Token) CreateFT(tokenCreation *CreateFungibleTokenDTO) (*Token, error) {
	tokenParams, err := tokenCreation.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid token creation params: %s", err)
	}

	tokenCreateTransaction := hedera.NewTokenCreateTransaction().
		SetTokenType(tokenParams.Type).
		SetTokenName(tokenParams.Name).
		SetTokenSymbol(tokenParams.Symbol).
		SetTreasuryAccountID(tokenParams.TreasuryAccountID).
		SetDecimals(tokenParams.Decimals).
		SetInitialSupply(tokenParams.InitSupply).
		SetSupplyType(tokenParams.SupplyType).
		SetMaxSupply(tokenParams.MaxSupply).
		SetFreezeDefault(tokenParams.FreezeDefault)

	if tokenParams.AdminPublicKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetAdminKey(tokenParams.AdminPublicKey)
	}

	if tokenParams.KycKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetKycKey(tokenParams.KycKey)
	}

	if tokenParams.FreezeKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetFreezeKey(tokenParams.FreezeKey)
	}

	if tokenParams.SupplyKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetSupplyKey(tokenParams.SupplyKey)
	}

	if tokenParams.FeeScheduleKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetFeeScheduleKey(tokenParams.FeeScheduleKey)
	}

	if tokenParams.PauseKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetPauseKey(tokenParams.PauseKey)
	}

	if tokenParams.WipeKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetWipeKey(tokenParams.WipeKey)
	}

	// if len(tokenCreation.CustomFees) > 0 {
	// 	tokenCreateTransaction = tokenCreateTransaction.SetCustomFees(tokenCreation.CustomFees)
	// }

	if tokenParams.AutoRenewAccount != (hedera.AccountID{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetAutoRenewAccount(tokenParams.AutoRenewAccount)
	}

	if len(tokenCreation.Memo) > 0 {
		tokenCreateTransaction = tokenCreateTransaction.SetTokenMemo(tokenCreation.Memo)
	}

	if tokenCreation.ExpirationTime != (time.Time{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetExpirationTime(tokenCreation.ExpirationTime)
	}
	tokenCreateTransaction = tokenCreateTransaction.SetFreezeDefault(tokenCreation.FreezeDefault)

	transaction, err := tokenCreateTransaction.FreezeWith(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("freeze transaction failed: %s", err)
	}

	txResponse, err := transaction.
		Sign(tokenParams.AdminPrivateKey).
		Sign(tokenParams.TreasuryPrivateKey).
		Execute(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("execute transaction failed: %s", err)
	}

	receipt, err := txResponse.GetReceipt(t.gateway.GetClient())
	if err != nil {
		return nil, fmt.Errorf("get transaction receipt failed: %s", err)
	}

	return t.WithTokenID(*receipt.TokenID), nil
}

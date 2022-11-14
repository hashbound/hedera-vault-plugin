package token

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type CreateFungibleTokenDTO struct {
	Type               string         `json:"type" validate:"required,oneof=TOKEN_TYPE_FUNGIBLE_COMMON TOKEN_TYPE_NON_FUNGIBLE_UNIQUE"`
	Name               string         `json:"name" validate:"required,min=1,max=100"`
	Symbol             string         `json:"symbol" validate:"required,min=1,max=100"`
	Decimals           uint           `json:"decimal" validate:"min=0"`
	InitSupply         uint64         `json:"initSupply" validate:"min=0"`
	TreasuryAccountID  string         `json:"treasuryAccountID" validate:"required"`
	TreasuryPrivateKey string         `json:"treasuryPrivateKey"`
	AdminPrivateKey    key.PrivateKey `json:"adminPrivateKey"`
	KycKey             key.PublicKey  `json:"kycKey"`
	FreezeKey          key.PublicKey  `json:"freezeKey"`
	WipeKey            key.PublicKey  `json:"wipeKey"`
	SupplyKey          key.PublicKey  `json:"supplyKey"`
	FeeScheduleKey     key.PublicKey  `json:"feeScheduleKey"`
	PauseKey           key.PublicKey  `json:"pauseKey"`
	AutoRenewAccount   key.PublicKey  `json:"autoRenewAccount"`
	// CustomFees         []string       `json:"customFees"`
	MaxSupply      int64     `json:"maxSupply" validate:"min=0"`
	SupplyType     string    `json:"supplyType" validate:"oneof=TOKEN_SUPPLY_TYPE_INFINITE TOKEN_SUPPLY_TYPE_FINITE"`
	FreezeDefault  bool      `json:"freezeDefault"`
	ExpirationTime time.Time `json:"expirationTime"`
	Memo           string    `json:"memo"`
}

func (tokenCreation *CreateFungibleTokenDTO) validate() (*CreateFungibleToken, error) {
	validate := validator.New()
	err := validate.Struct(tokenCreation)
	if err != nil {
		return nil, fmt.Errorf("invalid token creation parameters: %s", err)
	}

	t := &CreateFungibleToken{}

	if tokenCreation.Type == hedera.TokenTypeFungibleCommon.String() {
		t.Type = hedera.TokenTypeFungibleCommon
	} else if tokenCreation.Type == hedera.TokenTypeNonFungibleUnique.String() {
		t.Type = hedera.TokenTypeNonFungibleUnique
	}
	t.Name = tokenCreation.Name
	t.Symbol = tokenCreation.Symbol
	t.Decimals = tokenCreation.Decimals

	if tokenCreation.Type == hedera.TokenTypeNonFungibleUnique.String() && tokenCreation.Decimals > 0 {
		return nil, fmt.Errorf("invalid supply for NFT token")
	}
	t.InitSupply = tokenCreation.InitSupply

	treasuryID, err := hedera.AccountIDFromString(tokenCreation.TreasuryAccountID)
	if err != nil {
		return nil, fmt.Errorf("invalid treasury accountID: %s", err)
	}
	t.TreasuryAccountID = treasuryID

	treasuryPrivateKey, err := hedera.PrivateKeyFromString(tokenCreation.TreasuryPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid Treasury Private Key: %s", err)
	}
	t.TreasuryPrivateKey = treasuryPrivateKey

	if tokenCreation.AdminPrivateKey != (key.PrivateKey{}) {
		adminKey, err := key.FromPrivateKey(tokenCreation.AdminPrivateKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Key: %s", err)
		}
		t.AdminPrivateKey = adminKey.PrivateKey
		t.AdminPublicKey = adminKey.PublicKey
	}

	if tokenCreation.KycKey != (key.PublicKey{}) {
		kycKey, err := key.FromPublicKey(tokenCreation.KycKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.KycKey = kycKey.PublicKey
	}

	if tokenCreation.FreezeKey != (key.PublicKey{}) {
		freezeKey, err := key.FromPublicKey(tokenCreation.FreezeKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.FreezeKey = freezeKey.PublicKey
	}

	if tokenCreation.SupplyKey != (key.PublicKey{}) {
		supplyKey, err := key.FromPublicKey(tokenCreation.SupplyKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.SupplyKey = supplyKey.PublicKey
	}

	if tokenCreation.FeeScheduleKey != (key.PublicKey{}) {
		feeScheduleKey, err := key.FromPublicKey(tokenCreation.FeeScheduleKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.FeeScheduleKey = feeScheduleKey.PublicKey
	}

	if tokenCreation.PauseKey != (key.PublicKey{}) {
		pauseKey, err := key.FromPublicKey(tokenCreation.PauseKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.PauseKey = pauseKey.PublicKey
	}

	if tokenCreation.WipeKey != (key.PublicKey{}) {
		wipeKey, err := key.FromPublicKey(tokenCreation.WipeKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.WipeKey = wipeKey.PublicKey
	}

	if tokenCreation.AutoRenewAccount != (key.PublicKey{}) {
		autoRenewalID, err := hedera.AccountIDFromString(tokenCreation.TreasuryAccountID)
		if err != nil {
			return nil, fmt.Errorf("invalid AutoRenewal accountID: %s", err)
		}
		t.AutoRenewAccount = autoRenewalID
	}

	if tokenCreation.SupplyType == hedera.TokenSupplyTypeFinite.String() {
		t.SupplyType = hedera.TokenSupplyTypeFinite
		t.MaxSupply = tokenCreation.MaxSupply
	} else if tokenCreation.SupplyType == hedera.TokenSupplyTypeInfinite.String() {
		if tokenCreation.MaxSupply > 0 {
			return nil, fmt.Errorf("invalid max supply with supply type")
		}
		t.SupplyType = hedera.TokenSupplyTypeInfinite
	}

	t.FreezeDefault = tokenCreation.FreezeDefault
	t.ExpirationTime = tokenCreation.ExpirationTime
	t.Memo = tokenCreation.Memo

	return t, nil
}

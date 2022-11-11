package token

import (
	"fmt"
	"time"

	"github.com/hashbound/hedera-vault-plugin/src/core/key"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type CreateFungibleTokenDTO struct {
	Type               string         `json:"type"`
	Name               string         `json:"name"`
	Symbol             string         `json:"symbol"`
	Decimals           uint           `json:"decimal"`
	InitSupply         uint64         `json:"initSupply"`
	TreasuryAccountID  string         `json:"treasuryAccountID"`
	TreasuryPrivateKey string         `json:"treasuryPrivateKey"`
	AdminPrivateKey    key.PrivateKey `json:"adminPrivateKey"`
	KycKey             string         `json:"kycKey"`
	FreezeKey          string         `json:"freezeKey"`
	WipeKey            string         `json:"wipeKey"`
	SupplyKey          string         `json:"supplyKey"`
	FeeScheduleKey     string         `json:"feeScheduleKey"`
	PauseKey           string         `json:"pauseKey"`
	// CustomFees         []string       `json:"customFees"`
	MaxSupply        int64     `json:"maxSupply"`
	SupplyType       string    `json:"supplyType"`
	FreezeDefault    bool      `json:"freezeDefault"`
	ExpirationTime   time.Time `json:"expirationTime"`
	AutoRenewAccount string    `json:"autoRenewAccount"`
	Memo             string    `json:"memo"`
}

func (tokenCreation *CreateFungibleTokenDTO) validate() (*CreateFungibleToken, error) {
	t := &CreateFungibleToken{}

	if tokenCreation.Type == hedera.TokenTypeFungibleCommon.String() {
		t.Type = hedera.TokenTypeFungibleCommon
	} else if tokenCreation.Type == hedera.TokenTypeNonFungibleUnique.String() {
		t.Type = hedera.TokenTypeNonFungibleUnique
	} else {
		return nil, fmt.Errorf("invalid token type")
	}

	if len(tokenCreation.Name) == 0 || len(tokenCreation.Name) > 100 {
		return nil, fmt.Errorf("invalid token Name")
	}
	t.Name = tokenCreation.Name

	if len(tokenCreation.Symbol) == 0 || len(tokenCreation.Symbol) > 100 {
		return nil, fmt.Errorf("invalid token Symbol")
	}
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

	if tokenCreation.KycKey != "" {
		kycKey, err := hedera.PublicKeyFromString(tokenCreation.KycKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.KycKey = kycKey
	}

	if tokenCreation.FreezeKey != "" {
		freezeKey, err := hedera.PublicKeyFromString(tokenCreation.FreezeKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.FreezeKey = freezeKey
	}

	if tokenCreation.SupplyKey != "" {
		supplyKey, err := hedera.PublicKeyFromString(tokenCreation.SupplyKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.SupplyKey = supplyKey
	}

	if tokenCreation.FeeScheduleKey != "" {
		feeScheduleKey, err := hedera.PublicKeyFromString(tokenCreation.FeeScheduleKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.FeeScheduleKey = feeScheduleKey
	}

	if tokenCreation.PauseKey != "" {
		pauseKey, err := hedera.PublicKeyFromString(tokenCreation.PauseKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.PauseKey = pauseKey
	}

	if tokenCreation.WipeKey != "" {
		wipeKey, err := hedera.PublicKeyFromString(tokenCreation.WipeKey)
		if err != nil {
			return nil, fmt.Errorf("invalid Admin Public Key: %s", err)
		}
		t.WipeKey = wipeKey
	}

	if tokenCreation.AutoRenewAccount != "" {
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
			return nil, fmt.Errorf("invalid supply type with max supply")
		}
		t.SupplyType = hedera.TokenSupplyTypeInfinite
	} else {
		return nil, fmt.Errorf("invalid token supply type")
	}

	t.FreezeDefault = tokenCreation.FreezeDefault
	t.ExpirationTime = tokenCreation.ExpirationTime
	t.Memo = tokenCreation.Memo

	return t, nil
}
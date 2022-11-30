package store

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/account/entity"
)

type AccountStore struct {
	ctx         context.Context
	storage     logical.Storage
	clientToken string
}

const (
	PathIdentifier = "account"
)

func New(ctx context.Context, storage logical.Storage) *AccountStore {
	return &AccountStore{
		ctx:     ctx,
		storage: storage,
	}
}

func (as *AccountStore) WithClientToken(clientToken string) *AccountStore {
	as.clientToken = clientToken
	return as
}

func (as *AccountStore) Write(id string, account *entity.Account) error {
	account_bytes, err := account.ToBytes()
	if err != nil {
		return fmt.Errorf("encoding account to bytes failed: %s", err)
	}

	entry := &logical.StorageEntry{
		Key:      as.getPath(id),
		Value:    account_bytes,
		SealWrap: false,
	}
	if err := as.storage.Put(as.ctx, entry); err != nil {
		return fmt.Errorf("write account to storage failed: %s", err)
	}

	return nil
}

func (as *AccountStore) Read(id string) (*entity.Account, error) {
	entry, err := as.storage.Get(as.ctx, as.getPath(id))
	if err != nil {
		return nil, fmt.Errorf("fetch account from storage failed: %s", err)
	}
	if entry == nil {
		return nil, fmt.Errorf("account not found in storage")
	}

	account, err := entity.FromBytes(entry.Value)
	if err != nil {
		return nil, fmt.Errorf("parsing buffered account failed: %s", err)
	}

	return account, nil
}

func (as *AccountStore) List() ([]string, error) {
	entries, err := as.storage.List(as.ctx, as.getPath(""))
	if err != nil {
		return nil, fmt.Errorf("fetch accounts from storage filed: %s", err)
	}
	return entries, nil
}

func (as *AccountStore) Delete(id string) error {
	err := as.storage.Delete(as.ctx, as.getPath(id))
	if err != nil {
		return fmt.Errorf("delete account from storage failed: %s", err)
	}
	return nil
}

func (as *AccountStore) getPath(id string) string {
	return fmt.Sprintf("%s/%s/%s", as.clientToken, PathIdentifier, id)
}

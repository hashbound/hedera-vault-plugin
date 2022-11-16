package store

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/key/entity"
)

type KeyStore struct {
	ctx         context.Context
	storage     logical.Storage
	clientToken string
}

const (
	KeyPathIdentifier = "key"
)

func New(ctx context.Context, storage logical.Storage) *KeyStore {
	return &KeyStore{
		ctx:     ctx,
		storage: storage,
	}
}

func (ks *KeyStore) WithClientToken(clientToken string) *KeyStore {
	ks.clientToken = clientToken
	return ks
}

func (ks *KeyStore) Write(id string, key *entity.Key) error {
	key_bytes, err := key.ToBytes()
	if err != nil {
		return fmt.Errorf("encoding key to bytes failed: %s", err)
	}

	entry := &logical.StorageEntry{
		Key:      ks.getKeyPath(id),
		Value:    key_bytes,
		SealWrap: false,
	}
	if err := ks.storage.Put(ks.ctx, entry); err != nil {
		return fmt.Errorf("write key to storage failed: %s", err)
	}

	return nil
}

func (ks *KeyStore) Read(id string) (*entity.Key, error) {
	entry, err := ks.storage.Get(ks.ctx, ks.getKeyPath(id))
	if err != nil {
		return nil, fmt.Errorf("fetch key from storage failed: %s", err)
	}
	if entry == nil {
		return nil, fmt.Errorf("key not found in storage")
	}

	key, err := entity.FromBytes(entry.Value)
	if err != nil {
		return nil, fmt.Errorf("parsing buffered key failed: %s", err)
	}

	return key, nil
}

func (ks *KeyStore) List() ([]string, error) {
	entries, err := ks.storage.List(ks.ctx, ks.getKeyPath(""))
	if err != nil {
		return nil, fmt.Errorf("fetch keys from storage filed: %s", err)
	}
	return entries, nil
}

func (ks *KeyStore) Delete(id string) error {
	err := ks.storage.Delete(ks.ctx, ks.getKeyPath(id))
	if err != nil {
		return fmt.Errorf("delete key from storage failed: %s", err)
	}
	return nil
}

func (ks *KeyStore) getKeyPath(id string) string {
	return fmt.Sprintf("%s/%s/%s", ks.clientToken, KeyPathIdentifier, id)
}

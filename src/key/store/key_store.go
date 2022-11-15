package store

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
)

type KeyStore struct {
	ctx         context.Context
	storage     logical.Storage
	clientToken string
}

func New(ctx context.Context, storage logical.Storage) *KeyStore {
	return &KeyStore{
		ctx: ctx,
		storage: storage,
	}
}

func (ks *KeyStore) WithClientToken(clientToken string) *KeyStore {
	ks.clientToken = clientToken
	return ks
}
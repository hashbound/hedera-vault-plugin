package service

import (
	"context"

	"github.com/hashbound/hedera-vault-plugin/src/key/store"
	"github.com/hashicorp/vault/sdk/logical"
)

type KeyService struct {
	storage *store.KeyStore
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *KeyService {
	return &KeyService{
		storage: store.New(ctx, storage).WithClientToken(clientToken),
	}
}

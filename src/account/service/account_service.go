package service

import (
	"context"

	"github.com/hashbound/hedera-vault-plugin/src/account/store"
	"github.com/hashicorp/vault/sdk/logical"

	key_service "github.com/hashbound/hedera-vault-plugin/src/key/service"
)

type AccountService struct {
	storage *store.AccountStore
	k_svc   *key_service.KeyService
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *AccountService {
	return &AccountService{
		storage: store.New(ctx, storage).WithClientToken(clientToken),
		k_svc:   key_service.New(ctx, storage, clientToken),
	}
}

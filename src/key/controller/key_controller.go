package controller

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/key/service"
)

type KeyController struct {
	service *service.KeyService
}

func New(ctx context.Context, req *logical.Request) *KeyController {
	return &KeyController{
		service.New(ctx, req.Storage, req.ClientToken),
	}
}

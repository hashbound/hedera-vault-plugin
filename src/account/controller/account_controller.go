package controller

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/account/service"
)

type AccountController struct {
	service *service.AccountService
}

func New(ctx context.Context, req *logical.Request) *AccountController {
	return &AccountController{
		service.New(ctx, req.Storage, req.ClientToken),
	}
}

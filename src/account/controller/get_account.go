package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/formatters"
	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func Get(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	getAccountDTO := &dto.GetAccountDTO{
		ID:         data.Get("id").(string),
	}

	ac := New(ctx, req)
	account, err := ac.service.GetAccount(getAccountDTO)
	if err != nil {
		return nil, fmt.Errorf("get account failed")
	}

	return &logical.Response{
		Data: formatters.FormatResponse(account),
	}, nil
}

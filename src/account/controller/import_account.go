package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/formatters"
	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func Import(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	importAccountDTO := &dto.ImportAccountDTO{
		ID:         data.Get("id").(string),
		AccountID: data.Get("accountID").(string),
		KeyID: data.Get("keyID").(string),
	}

	ac := New(ctx, req)
	account, err := ac.service.ImportAccount(importAccountDTO)
	if err != nil {
		return nil, fmt.Errorf("import account failed")
	}

	return &logical.Response{
		Data: formatters.FormatResponse(account),
	}, nil
}

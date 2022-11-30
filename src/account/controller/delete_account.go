package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func Delete(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	deleteAccountDTO := &dto.DeleteAccountDTO{
		ID: data.Get("id").(string),
	}

	ac := New(ctx, req)
	err := ac.service.DeleteAccount(deleteAccountDTO)
	if err != nil {
		return nil, fmt.Errorf("remove account failed: %s", err)
	}

	return nil, nil
}
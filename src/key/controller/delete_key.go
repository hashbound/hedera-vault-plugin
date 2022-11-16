package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func Delete(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	deleteKeyDTO := &dto.DeleteKeyDTO{
		ID: data.Get("id").(string),
	}

	kc := New(ctx, req)
	return nil, kc.service.Delete(deleteKeyDTO)
}

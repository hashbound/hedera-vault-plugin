package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/core/formatters"
	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func GetKey(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	getKeyDTO := &dto.GetKeyDTO{
		ID: data.Get("id").(string),
	}

	kc := New(ctx, req)
	key, err := kc.service.GetKey(getKeyDTO)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: formatters.FormatResponse(key),
	}, nil
}

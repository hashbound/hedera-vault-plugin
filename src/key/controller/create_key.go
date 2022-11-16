package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/core/formatters"
	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
)

func Create(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	createKeyDTO := &dto.CreateKeyDTO{
		ID:        data.Get("id").(string),
		Algorithm: data.Get("algo").(string),
		Curve:     data.Get("curve").(string),
	}

	kc := New(ctx, req)
	key, err := kc.service.Create(createKeyDTO)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: formatters.FormatResponse(key),
	}, nil
}

package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/core/formatters"
	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
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

	importKeyDTO := &dto.ImportKeyDTO{
		ID:         data.Get("id").(string),
		PrivateKey: data.Get("privateKey").(string),
		Algorithm:  data.Get("algo").(string),
		Curve:      data.Get("curve").(string),
	}

	kc := New(ctx, req)
	key_vault, err := kc.service.ImportKey(importKeyDTO)
	if err != nil {
		return nil, fmt.Errorf("import key failed")
	}

	return &logical.Response{
		Data: formatters.FormatResponse(key_vault),
	}, nil
}

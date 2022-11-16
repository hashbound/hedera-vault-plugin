package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/key/dto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func Sign(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	signMessageDTO := &dto.SignMessageDTO{
		ID:      data.Get("id").(string),
		Message: data.Get("message").(string),
	}

	kc := New(ctx, req)
	signature, err := kc.service.Sign(signMessageDTO)
	if err != nil {
		return nil, fmt.Errorf("sign message failed: %s", err)
	}

	resp := make(map[string]interface{})
	resp["signature"] = signature

	return &logical.Response{
		Data: resp,
	}, nil
}

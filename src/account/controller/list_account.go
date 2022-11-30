package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func List(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	ac := New(ctx, req)
	accounts, err := ac.service.List()
	if err != nil {
		return nil, fmt.Errorf("list accounts failed: %s", err)
	}

	return logical.ListResponse(accounts), nil
}

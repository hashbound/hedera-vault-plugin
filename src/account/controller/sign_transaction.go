package controller

import (
	"context"
	"fmt"

	"github.com/hashbound/hedera-vault-plugin/src/account/dto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func SignTransaction(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	signTransaactionDTO := &dto.SignTransactionDTO{
		AccountID:   data.Get("accountId").(string),
		Transaction: data.Get("transaction").(string),
	}

	ac := New(ctx, req)
	signedTransaction, err := ac.service.SignTransaction(signTransaactionDTO)
	if err != nil {
		return nil, fmt.Errorf("sign transaction failed: %s", err)
	}

	response := make(map[string]interface{})
	response["signed"] = string(signedTransaction)

	return &logical.Response{
		Data: response,
	}, nil
}

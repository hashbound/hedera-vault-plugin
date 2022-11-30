package path

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/account/controller"
)

type AccountPaths struct {
}

func NewKeyPaths() *AccountPaths {
	return &AccountPaths{}
}

func (ap *AccountPaths) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			ap.pathAccounts(),
			ap.pathImportAccounts(),
			ap.pathSignTransaction(),
		},
	)
}

func (kp *AccountPaths) pathAccounts() *framework.Path {
	return &framework.Path{
		Pattern: "accounts/?",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: &framework.PathOperation{
				Callback: controller.Get,
			},
			logical.ListOperation: &framework.PathOperation{
				Callback: controller.List,
			},
			logical.DeleteOperation: &framework.PathOperation{
				Callback: controller.Delete,
			},
		},

		ExistenceCheck: kp.handleExistenceCheck,
	}
}

func (kp *AccountPaths) pathImportAccounts() *framework.Path {
	return &framework.Path{
		Pattern: "accounts/import",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"accountId": {
				Type:     framework.TypeString,
				Required: true,
			},
			"keyId": {
				Type:     framework.TypeString,
				Required: true,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: controller.Import,
			},
		},

		ExistenceCheck: kp.handleExistenceCheck,
	}
}

func (kp *AccountPaths) pathSignTransaction() *framework.Path {
	return &framework.Path{
		Pattern: "accounts/sign_transaction",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"transaction": {
				Type:     framework.TypeString,
				Required: true,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: controller.SignTransaction,
			},
		},

		ExistenceCheck: kp.handleExistenceCheck,
	}
}

func (kp *AccountPaths) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, fmt.Errorf("existence check failed: %s", err)
	}

	return out != nil, nil
}

package path

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashbound/hedera-vault-plugin/src/key/controller"
)

type KeyPaths struct {
}

func NewKeyPaths() *KeyPaths {
	return &KeyPaths{}
}

func (kp *KeyPaths) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			kp.pathKeys(),
			kp.pathImportKeys(),
			kp.pathSign(),
		},
	)
}

func (kp *KeyPaths) pathKeys() *framework.Path {
	return &framework.Path{
		Pattern: "keys/?",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"algo": {
				Type:     framework.TypeString,
				Required: true,
			},
			"curve": {
				Type: framework.TypeString,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: controller.Create,
			},
			logical.ReadOperation: &framework.PathOperation{
				Callback: controller.GetKey,
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

func (kp *KeyPaths) pathImportKeys() *framework.Path {
	return &framework.Path{
		Pattern: "keys/import",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"privateKey": {
				Type:     framework.TypeString,
				Required: true,
			},
			"algo": {
				Type:     framework.TypeString,
				Required: true,
			},
			"curve": {
				Type: framework.TypeString,
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

func (kp *KeyPaths) pathSign() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("keys/%s/sign", framework.GenericNameRegex("id")),

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"message": {
				Type:     framework.TypeString,
				Required: true,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: controller.Sign,
			},
		},

		ExistenceCheck: kp.handleExistenceCheck,
	}
}

func (kp *KeyPaths) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, fmt.Errorf("existence check failed: %s", err)
	}

	return out != nil, nil
}

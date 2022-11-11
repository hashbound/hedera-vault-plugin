package client

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (c *Client) WithOperator(operatorID, operatorKey string) *Client {
	c.operatorID = operatorID
	c.operatorKey = operatorKey

	return c
}

func (c *Client) SetOperator() (*Client, error) {
	operatorID, err := hedera.AccountIDFromString(c.operatorID)
	if err != nil {
		return nil, fmt.Errorf("invalid operatorID: %s", err)
	}

	operatorKey, err := hedera.PrivateKeyFromString(c.operatorKey)
	if err != nil {
		return nil, fmt.Errorf("invalid operator key: %s", err)
	}
	
	c.client = c.client.SetOperator(operatorID, operatorKey)
	
	return c, nil
}
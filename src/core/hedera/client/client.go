package client

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Client struct {
	client      *hedera.Client
	operatorID  string
	operatorKey string
}

func New() *Client {
	return &Client{}
}

func (c *Client) ClientFromConfig(cfg *ClientConfig) (*Client, error) {
	accountID, err := hedera.AccountIDFromString(cfg.NetworkNodeAccountID)
	if err != nil {
		return nil, fmt.Errorf("invalid network account id: %s", err)
	}

	node := make(map[string]hedera.AccountID, 1)
	node[cfg.NetworkNodeAddress] = accountID
	c.client = hedera.ClientForNetwork(node)

	if c.operatorID != "" && c.operatorKey != "" {
		c, err = c.SetOperator()
		if err != nil {
			return nil, fmt.Errorf("unable to set operator: %s", err)
		}
	}

	return c, nil
}

func (c *Client) ClientFromNetworkName(network string) (*Client, error) {
	client, err := hedera.ClientForName(network)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare client: %s", err)
	}
	c.client = client

	if c.operatorID != "" && c.operatorKey != "" {
		c, err = c.SetOperator()
		if err != nil {
			return nil, fmt.Errorf("unable to set operator: %s", err)
		}
	}

	return c, nil
}

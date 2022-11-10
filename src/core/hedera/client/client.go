package client

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

const (
	NETWORK_MAINNET = "mainnet"
	NETWORK_TESTNET = "testnet"
)

type ClientConfig struct {
	NetworkNodeAddress   string
	NetworkNodeAccountID string
}

func DefaultLocalTestNetConfig() *ClientConfig {
	return &ClientConfig{
		NetworkNodeAddress:   "127.0.0.1:50211",
		NetworkNodeAccountID: "0.0.3",
	}
}

func ClientFromConfig(cfg *ClientConfig) (*hedera.Client, error) {
	accountID, err := hedera.AccountIDFromString(cfg.NetworkNodeAccountID)
	if err != nil {
		return nil, fmt.Errorf("invalid network account id: %s", err)
	}

	node := make(map[string]hedera.AccountID, 1)
	node[cfg.NetworkNodeAddress] = accountID

	client := hedera.ClientForNetwork(node)

	return client, nil
}

func ClientFromNetworkName(network string) (*hedera.Client, error) {
	return hedera.ClientForName(network)
}

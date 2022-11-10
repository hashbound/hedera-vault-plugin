package client

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

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

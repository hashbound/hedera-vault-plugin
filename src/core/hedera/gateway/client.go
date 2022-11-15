package gateway

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (gw *Gateway) ClientFromConfig(cfg *ClientConfig) (*Gateway, error) {
	accountID, err := hedera.AccountIDFromString(cfg.NetworkNodeAccountID)
	if err != nil {
		return nil, fmt.Errorf("invalid network account id: %s", err)
	}

	node := make(map[string]hedera.AccountID, 1)
	node[cfg.NetworkNodeAddress] = accountID
	gw.client = hedera.ClientForNetwork(node)

	return gw, nil
}

func (gw *Gateway) ClientFromNetworkName(network string) (*Gateway, error) {
	client, err := hedera.ClientForName(network)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare client: %s", err)
	}
	gw.client = client

	return gw, nil
}

package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientFromConfig(t *testing.T) {
	client, err := ClientFromConfig(DefaultLocalTestNetConfig())
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}
	
	fmt.Printf("client: %v", client.GetLedgerID())
}

func TestClientTestNet(t *testing.T) {
	client, err := ClientFromNetworkName(NETWORK_TESTNET)
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}

	fmt.Printf("client: %v", client.GetLedgerID())
	assert.Equal(t, true, client.GetLedgerID().IsTestnet(), fmt.Sprintf("expected: %v\nreceived:%v", true, client.GetLedgerID().IsTestnet()))
}

func TestClientMainNet(t *testing.T) {
	client, err := ClientFromNetworkName(NETWORK_MAINNET)
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}

	fmt.Printf("client: %v", client.GetLedgerID())
	assert.Equal(t, true, client.GetLedgerID().IsMainnet(), fmt.Sprintf("expected: %v\nreceived:%v", true, client.GetLedgerID().IsTestnet()))
}
package gateway

import (
	"fmt"
	"testing"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

func TestClientFromConfig(t *testing.T) {
	g := New()
	g, err := g.ClientFromConfig(DefaultLocalTestNetConfig())
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}

	fmt.Printf("client: %v", g.client.GetLedgerID())
}

func TestClientTestNet(t *testing.T) {
	g := New()
	g, err := g.ClientFromNetworkName(string(hedera.NetworkNameTestnet))
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}

	fmt.Printf("client: %v", g.client.GetLedgerID())
	assert.Equal(t, true, g.client.GetLedgerID().IsTestnet(), fmt.Sprintf("expected: %v\nreceived:%v", true, g.client.GetLedgerID().IsTestnet()))
}

func TestClientMainNet(t *testing.T) {
	g := New()
	g, err := g.ClientFromNetworkName(string(hedera.NetworkNameMainnet))
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}

	fmt.Printf("client: %v", g.client.GetLedgerID())
	assert.Equal(t, true, g.client.GetLedgerID().IsMainnet(), fmt.Sprintf("expected: %v\nreceived:%v", true, g.client.GetLedgerID().IsTestnet()))
}

func TestClientFromConfigWithOperator(t *testing.T) {
	g := New()
	g, err := g.
		WithOperator("0.0.2", "302e020100300506032b65700422042091132178e72057a1d7528025956fe39b0b847f200ab59b2fdd367017f3087137").
		ClientFromConfig(DefaultLocalTestNetConfig())
	if err != nil {
		t.Fatalf("unable to create client: %s", err)
	}

	fmt.Printf("client: %v", g.client.GetLedgerID())
}
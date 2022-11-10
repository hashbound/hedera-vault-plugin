package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientConfig(t *testing.T) {
	networkAccountID := "0.0.3"
	networkAddress := "127.0.0.1:50211"

	cc := NewClientConfig().WithNetworkAccountID(networkAccountID).WithNetworkAddress(networkAddress)
	
	assert.Equal(t, networkAccountID, cc.NetworkNodeAccountID, fmt.Sprintf("expected: %s\n,received: %s", networkAccountID, cc.NetworkNodeAccountID))
	assert.Equal(t, networkAddress, cc.NetworkNodeAddress, fmt.Sprintf("expected: %s\n,received: %s", networkAddress, cc.NetworkNodeAddress))
}
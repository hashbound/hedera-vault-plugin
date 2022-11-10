package client

type ClientConfig struct {
	NetworkNodeAddress   string
	NetworkNodeAccountID string
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{}
}

func DefaultLocalTestNetConfig() *ClientConfig {
	return &ClientConfig{
		NetworkNodeAddress:   "127.0.0.1:50211",
		NetworkNodeAccountID: "0.0.3",
	}
}

func (c *ClientConfig) WithNetworkAddress(node string) *ClientConfig {
	c.NetworkNodeAddress = node
	return c
}

func (c *ClientConfig) WithNetworkAccountID(accountID string) *ClientConfig {
	c.NetworkNodeAccountID = accountID
	return c
}
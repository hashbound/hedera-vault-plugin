package client

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
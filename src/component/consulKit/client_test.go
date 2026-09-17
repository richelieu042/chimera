package consulKit

import (
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestNewClient(t *testing.T) {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	config.Scheme = "http"

	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	client = client
}

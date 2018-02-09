package openei

import "github.com/buskersguidetotheuniverse.org/types"

// documentation:
//https://openei.org/services/doc/rest/util_rates/?version=3
// API:  https://api.openei.org/utility_rates?parameters
const apiBaseUrl = "https://api.openei.org/utility_rates"

type Client struct {
	ApiKey        string
	DefaultParams map[string]string
}

func NewClient(apiKey string) Client {
	return Client{
		ApiKey:        apiKey,
		DefaultParams: map[string]string{"version": "latest", "format": "json", "api_key": apiKey, "limit": "10"},
	}
}

func (client Client) CurrentEnergyPrices(location types.Geometry) (interface{}, error) {
	var err error

	return nil, err
}

package openei

import (
	"encoding/json"
	"fmt"
	"github.com/buskersguidetotheuniverse.org/net"
	"github.com/buskersguidetotheuniverse.org/types"
)

//https://api.openei.org/utility_rates?api_key=xxx&
// version=latest&
// format=json&
// limit=1&
// sector=Residential&
// approved=true&
// lat=44.9778&lon=-93.2650&
// detail=full

// documentation:
//https://openei.org/services/doc/rest/util_rates/?version=3
// API:  https://api.openei.org/utility_rates?parameters
const apiBaseUrl = "https://api.openei.org/utility_rates?api_key="

type Client struct {
	ApiKey        string
	DefaultParams map[string]string
}

func NewClient(apiKey string) Client {
	return Client{
		ApiKey:        apiKey,
		DefaultParams: map[string]string{"version": "latest", "format": "json", "limit": "1", "sector": "Residential", "approved": "true", "detail": "full"},
	}
}

func (client Client) CurrentEnergyPrices(location *types.Geometry) (types.OpenEIResponse, error) {
	var err error

	url := apiBaseUrl + client.ApiKey
	for k, v := range client.DefaultParams {
		url += fmt.Sprintf(url, "&%v=%v", k, v)
	}

	// FIXME: Looks like we're building this wrong.
	url += fmt.Sprintf("&lat=%v&lon=%v", location.Coordinates[0], location.Coordinates[1])

	fmt.Printf("energy api url: %v\n", url)

	body, err := net.ReadFromUrl(url)
	fmt.Printf("response: %v\n", body)

	var response types.OpenEIResponse
	err = json.Unmarshal(body, &response)

	return response, err
}

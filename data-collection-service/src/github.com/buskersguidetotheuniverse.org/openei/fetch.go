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
		DefaultParams: map[string]string{"version": "latest", "format": "json", "limit": "1", "sector": "Commercial", "approved": "true", "detail": "full"},
	}
}

func (client Client) CurrentEnergyPrices(location *types.Geometry) ([][]types.EnergyRate, error) {
	var err error

	url := apiBaseUrl + client.ApiKey
	var queryString string
	for k, v := range client.DefaultParams {
		queryString += fmt.Sprintf("&%v=%v", k, v)
	}

	queryString += fmt.Sprintf("&lat=%v&lon=%v", location.Coordinates[0], location.Coordinates[1])
	url = url + queryString

	//fmt.Printf("energy api url: %v\n", url)

	body, err := net.ReadFromUrl(url)
	fmt.Printf("raw response: %v\n", string(body))

	var response types.OpenEIResponse
	err = json.Unmarshal(body, &response)

	fmt.Printf("parsed response: %v\n", response)

	return response.Items[0].EnergyRateStructure, err
}

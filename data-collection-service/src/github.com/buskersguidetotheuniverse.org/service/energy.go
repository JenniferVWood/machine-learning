package service

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/openei"
	"github.com/buskersguidetotheuniverse.org/types"
	"sync"
)

type EnergyService struct {
	ApiKey string
	WG     *sync.WaitGroup
}

func NewEnergyService(apiKey string, wg *sync.WaitGroup) EnergyService {
	return EnergyService{
		ApiKey: apiKey,
		WG:     wg,
	}
}

func (service EnergyService) ProcessLocation(location *types.Geometry) error {
	client := openei.NewClient(service.ApiKey)

	service.WG.Add(1)
	// for now, this is atomic.  next weekend that'll be different.
	// get result from energy.Fetch
	rates, _ := client.CurrentEnergyPrices(location)

	fmt.Printf("Rates:", rates)

	// persist to hadoop
	err := hbase.SaveEnergyPrices(&rates)

	service.WG.Done()
	return err
}

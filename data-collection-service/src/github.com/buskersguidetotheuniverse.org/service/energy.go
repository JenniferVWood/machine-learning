package service

import (
	"fmt"
	//"github.com/buskersguidetotheuniverse.org/hbase"
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
	var err error
	client := openei.NewClient(service.ApiKey)

	service.WG.Add(1)
	// for now, this is atomic.  next weekend that'll be different.
	// get result from energy.Fetch
	response, _ := client.CurrentEnergyPrices(location)

	// persist to hadoop
	if len(response) > 0 {
		for index, _ := range response[0] {
			response[0][index].Geometry = *location
		}
		fmt.Printf("Rates:%v\n", response[0])
		err = hbase.SaveEnergyPrices(&response[0])
	}
	service.WG.Done()
	return err
}

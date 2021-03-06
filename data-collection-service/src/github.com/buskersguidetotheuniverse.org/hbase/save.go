package hbase

import (
	//"github.com/tsuna/gohbase"
	"log"

	"context"
	"encoding/json"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/buskersguidetotheuniverse.org/openei"
	"github.com/buskersguidetotheuniverse.org/types"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

/*
	save data from NOAA to the local Cassandra instance
*/

func SaveObservation(observation *types.CurrentConditionsResponse) error {

	key, err := noaa.MakeKeyFromTimeStamp(noaa.StationShortForm(observation.Props.Station), observation.Props.Timestamp)
	if err != nil {
		log.Fatalf("Error making key: %v", err)
	}

	log.Printf("generated key: %v\n", key)

	client := gohbase.NewClient("localhost")
	//defer client.Close()  -- causes error??  "Client is dead"

	data, err := json.Marshal(observation.Props)
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)

	}

	family := map[string]map[string][]byte{"data": {"data": data}}

	putRequest, err := hrpc.NewPutStr(context.Background(), "observations", key, family)
	rsp, err := client.Put(putRequest)

	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}

	log.Printf("Inserted row.  Response from server: %v", rsp)

	return err
}

func SaveEnergyPrices(rates *[]types.EnergyRate) error {
	client := gohbase.NewClient("localhost")

	for _, rate := range *rates {

		key := openei.MakeKeyFromTimeStampAndGeo(rate.Geometry)

		data, err := json.Marshal(rate)

		family := map[string]map[string][]byte{"rates": {"rates": data}}

		putRequest, err := hrpc.NewPutStr(context.Background(), "energy", key, family)
		_, err = client.Put(putRequest)

		if err != nil {
			log.Fatalf("Error inserting data: %v", err)
		}

		log.Printf("Inserted row.")

	}

	return nil
}

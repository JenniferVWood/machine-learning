package hbase

import (
//"github.com/tsuna/gohbase"
	"log"

	"github.com/buskersguidetotheuniverse.org/types"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"context"
	"encoding/json"
)

/*
	save data from NOAA to the local Cassandra instance
*/

func SaveObservation(observation *types.CurrentConditionsResponse) {

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

	family :=  map[string]map[string][]byte{"data": {"data": data}}

	putRequest, err := hrpc.NewPutStr(context.Background(), "observations", key, family)
	rsp, err := client.Put(putRequest)

	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}

	log.Printf("Inserted row.  Response from server: %v", rsp)
}

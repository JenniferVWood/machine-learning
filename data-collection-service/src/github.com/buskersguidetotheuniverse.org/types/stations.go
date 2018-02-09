package types

type StationsResponse struct {
	Context             interface{} `json:"@Context"`
	Type                string      `json:"type"`
	Features            []Feature   `json:"features"`
	ObservationStations []string    `json:"observationStations"`
}

type Feature struct {
	Id            string            `json:"id"`
	Type          string            `json:"type"`
	Geometry      Geometry          `json:"geometry"`
	Properties    StationProperties `json:"properties"`
	QueryLocation Geometry          `json:"QueryLocation"` // populated locally
}

type StationProperties struct {
	Id                string   `json:"@id"`
	Type              string   `json:"@type"`
	Elevation         Property `json:"elevation"`
	StationIdentifier string   `json:"stationIdentifier"`
	Name              string   `json:"name"`
	TimeZone          string   `json:"timeZone"`
}

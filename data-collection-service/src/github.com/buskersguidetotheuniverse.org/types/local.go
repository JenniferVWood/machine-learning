package types

// contains types used internally

type ProcessStationRequest struct {
	StationProperties StationProperties // populated when operating on a lat-long
	StationId         string            // when stationId is passed in explicitly
	QueryLocation     Geometry
	PrintWeather      bool
}

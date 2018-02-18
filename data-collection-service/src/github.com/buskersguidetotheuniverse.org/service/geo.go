package service

import (
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"github.com/buskersguidetotheuniverse.org/types"
)

// TODO: make this a part of the Geometry type?
func GetDistance(from *types.Geometry, to *types.Geometry) types.Distance {

	lat1, lon1 := from.Coordinates[0], from.Coordinates[1]
	lat2, lon2 := to.Coordinates[0], to.Coordinates[1]

	// Create Ellipsoid object with WGS84-ellipsoid,
	// angle units are degrees, distance units are meter.
	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	// Calculate the distance and bearing from SFO to LAX.
	distance, bearing := geo1.To(lat1, lon1, lat2, lon2)

	return types.Distance{
		Bearing:     bearing,
		BearingUnit: "DEG",
		Range:       distance,
		RangeUnit:   "kilometers", // I think there's a bug in golang-elipsoid, that confuses M with KM.
	}
}

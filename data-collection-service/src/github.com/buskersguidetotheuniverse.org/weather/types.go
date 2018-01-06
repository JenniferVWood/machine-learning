package weather

/*
	structs to model data as it's received from NOA APIs,
	and to model data to be inserted into tables in the 'weather' Cassandra namespace.
 */


 type Property struct {
 	name string
 	value float32
 	unitCode string
 	qualityControl string
 }

 type Sample struct {
 	stationId string
 	timestamp int
 	properties []Property
 }
package types

/*
	structs to model data as it's received from NOA APIs,
	and to model data to be inserted into tables in the 'weather' Cassandra namespace.
*/

type CurrentConditionsResponse struct {
	Context  interface{} `json:"@Context"`
	Id       string      `json:"id"`
	Type     string      `json:"type"` // lowercase type is a reserved word...
	Geometry interface{} `json:"geometry"`
	Props    Properties  `json:"properties"`
}

type Properties struct {
	Id                        string       `json:"id"`
	Type                      string       `json:"type"`
	Elevation                 Property     `json:"elevation"`
	Station                   string       `json:"station"`
	Timestamp                 string       `json:"timestamp"`
	RawMessage                string       `json:"rawMessage"`
	TextDescription           string       `json:"textDescription"`
	Icon                      string       `json:"icon"`
	PresentWeather            interface{}  `json:"presentWeather"`
	Temperature               Property     `json:"temperature"`
	Dewpoint                  Property     `json:"dewPoint"`
	WindDirection             Property     `json:"windDirection"`
	WindSpeed                 Property     `json:"windSpeed"`
	WindGust                  Property     `json:"windGust"`
	BarometricPressure        Property     `json:"barometricPressure"`
	SeaLevelPressure          Property     `json:"seaLevelPressure"`
	Visibility                Property     `json:"visibility"`
	MaxTemperatureLast24Hours Property     `json:"maxTemperatureLast24Hours"`
	PrecipitationLastHour     Property     `json:"precipitationLastHour"`
	PrecipitationLast3Hours   Property     `json:"precipitationLast3Hours"`
	PrecipitationLast6Hours   Property     `json:"precipitationLast6Hours"`
	RelativeHumidity          Property     `json:"relativeHumidity"`
	WindChill                 Property     `json:"windChill"`
	HeatIndex                 Property     `json:"heatIndex"`
	CloudLayers               []CloudLayer `json:"cloudLayers"`
}

type Property struct {
	Value          float32 `json:"value"`
	UnitCode       string  `json:"unitCode"`
	QualityControl string  `json:"qualityControl"`
}

type CloudLayer struct {
	Base   Property `json:"base"`
	Amount string   `json:"Amount"`
}

type Geometry struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}

package types

type EnergyRateStructure struct {
	Tier       int     `json:"tier"`
	Rate       float32 `json:"rate"`
	Adjustment float32 `json:"adj"`
	Unit       string  `json:"unit"`
}

type OpenEIResponseItem struct {
	RateStructure EnergyRateStructure `json:"energyratestructure"`
}

type OpenEIResponse struct {
	Items []OpenEIResponseItem `json:"items"`
}

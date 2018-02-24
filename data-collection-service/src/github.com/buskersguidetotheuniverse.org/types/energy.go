package types

type EnergyRate struct {
	Rate       float32 `json:"rate"`
	Adjustment float32 `json:"adj"`
	Unit       string  `json:"unit"`
	Geometry   Geometry
}

type OpenEIResponseItem struct {
	Label                 string         `json:"label"`
	Uri                   string         `json:"uri"`
	Sector                string         `json:"sector"`
	EnergyWeekendSchedule interface{}    `json:"energyweekendschedule"`
	EnergyRateStructure   [][]EnergyRate `json:"energyratestructure"`
}

type OpenEIResponse struct {
	Items []OpenEIResponseItem `json:"items"`
}

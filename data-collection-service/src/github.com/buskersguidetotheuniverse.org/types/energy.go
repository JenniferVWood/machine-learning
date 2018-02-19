package types

type EnergyRateStructure struct {
	Tier       int     //`json:"tier"`
	Rate       float32 //`json:"rate"`
	Adjustment float32 //`json:"adj"`
	Unit       string  //`json:"unit"
}

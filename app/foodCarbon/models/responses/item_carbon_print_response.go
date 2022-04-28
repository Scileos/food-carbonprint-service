package responses

//ItemCarbonPrint ... ItemCarbonPrint response model
type ItemCarbonPrint struct {
	Name        string `json:"name"`
	CarbonPrint int64  `josn:"carbonPrint"`
}

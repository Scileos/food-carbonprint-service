package config

//FoodCarbonConfig ... FoodCarbonConfig type
type FoodCarbonConfig struct {
	CarbonValues string
}

//NewFoodCarbonConfig ... Constructor method for FoodCarbonConfig
func NewFoodCarbonConfig() *FoodCarbonConfig {
	config := "{\"beef\": 50,\"chicken\": 20,\"apple\": 1}"
	return &FoodCarbonConfig{CarbonValues: config}
}

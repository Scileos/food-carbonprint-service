package adapters

import (
	"github.com/Scileos/food-carbonprint-service/app/foodCarbon/models/shared"
	"github.com/Scileos/food-carbonprint-service/config"
	"github.com/tidwall/gjson"
)

//IFoodCarbonRepository ... FoodCarbonRepository interface
type IFoodCarbonRepository interface {
	GetCarbonPrintForBasket(items []string) (BasketCarbonPrintModel, error)
}

//FoodCarbonRepository ... FoodCarbonRepository type
type FoodCarbonRepository struct {
	FoodCarbonPrintConfig *config.FoodCarbonConfig
}

//BasketCarbonPrintModel ... Data access layer model for basket carbon print
type BasketCarbonPrintModel struct {
	Items []shared.ItemCarbonPrint
	Total int64
}

//NewFoodCarbonRepository ... Constructor method for NewFoodCarbonRepository
func NewFoodCarbonRepository(foodCarbonPrintConfig *config.FoodCarbonConfig) *FoodCarbonRepository {
	return &FoodCarbonRepository{FoodCarbonPrintConfig: foodCarbonPrintConfig}
}

//GetCarbonPrintForBasket ... Method to return a detailed carbon print model for given items
//Items not found will return -1 carbon value
func (f FoodCarbonRepository) GetCarbonPrintForBasket(items []string) (basketCarbonPrint BasketCarbonPrintModel, err error) {
	response := BasketCarbonPrintModel{}
	for _, item := range items {
		value := gjson.Get(f.FoodCarbonPrintConfig.CarbonValues, item)
		if value.Exists() {
			response.Items = append(response.Items, shared.ItemCarbonPrint{Name: item, CarbonPrint: value.Int()})
			response.Total = response.Total + value.Int()
		} else {
			response.Items = append(response.Items, shared.ItemCarbonPrint{Name: item, CarbonPrint: -1})
		}
	}

	return response, nil
}

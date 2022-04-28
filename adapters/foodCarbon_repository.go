package adapters

import (
	"github.com/Scileos/food-carbonprint-service/app/foodCarbon/models/shared"
	"github.com/Scileos/food-carbonprint-service/config"
	"github.com/tidwall/gjson"
)

//IFoodCarbonRepository ... FoodCarbonRepository interface
type IFoodCarbonRepository interface {
	GetCarbonPrintForBasket(items []string) (BasketCarbonPrintModel, error)
	GetCarbonPrintForItem(itemName string) (shared.ItemCarbonPrint, error)
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
	for _, item := range items {
		value := gjson.Get(f.FoodCarbonPrintConfig.CarbonValues, item)
		if value.Exists() {
			basketCarbonPrint.Items = append(basketCarbonPrint.Items, shared.ItemCarbonPrint{Name: item, CarbonPrint: value.Int()})
			basketCarbonPrint.Total = basketCarbonPrint.Total + value.Int()
		} else {
			basketCarbonPrint.Items = append(basketCarbonPrint.Items, shared.ItemCarbonPrint{Name: item, CarbonPrint: -1})
		}
	}

	return basketCarbonPrint, nil
}

//GetCarbonPrintForItem ... Method to return a detailed carbon print model for a given item
//Items not found will return -1 carbon value
func (f FoodCarbonRepository) GetCarbonPrintForItem(itemName string) (itemCarbonPrint shared.ItemCarbonPrint, err error) {
	itemCarbonPrint.Name = itemName

	value := gjson.Get(f.FoodCarbonPrintConfig.CarbonValues, itemName)
	if value.Exists() {
		itemCarbonPrint.CarbonPrint = value.Int()
	} else {
		itemCarbonPrint.CarbonPrint = -1
	}

	return itemCarbonPrint, nil
}

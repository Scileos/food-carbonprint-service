package responses

import (
	"github.com/Scileos/food-carbonprint-service/app/foodCarbon/models/shared"
)

//BasketCarbonPrint ... BasketCarbonPrint response model
type BasketCarbonPrint struct {
	Items []shared.ItemCarbonPrint `json:"items"`
	Total int64                    `josn:"total"`
}

package app

import (
	"github.com/Scileos/food-carbonprint-service/adapters"
	"github.com/Scileos/food-carbonprint-service/app/foodCarbon/models/requests"
	"github.com/Scileos/food-carbonprint-service/app/foodCarbon/models/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

//FoodCarbonService ... FoodCarbonService type
type FoodCarbonService struct {
	repo adapters.IFoodCarbonRepository
}

//NewFoodCarbonService ... Constructor method for FoodCarbonService
func NewFoodCarbonService(repo adapters.IFoodCarbonRepository) *FoodCarbonService {
	if repo == nil {
		panic("missing userRepository")
	}

	return &FoodCarbonService{repo: repo}
}

//GetCarbonPrintForBasket ... With an injected data access repo, get the carbon print values for a fiven set of items
func (fc FoodCarbonService) GetCarbonPrintForBasket(ctx *gin.Context) {
	var requestBody requests.BasketCarbonPrint
	validate := validator.New()

	err := ctx.BindJSON(&requestBody)
	err = validate.Struct(requestBody)

	if err != nil {
		ctx.JSON(400, "Invalid request")
		return
	}

	basketCarbonPrint, err := fc.repo.GetCarbonPrintForBasket(requestBody.Items)
	if err != nil {
		ctx.JSON(500, ctx.Error(err))
	} else {
		ctx.JSON(200, responses.BasketCarbonPrint{Items: basketCarbonPrint.Items, Total: basketCarbonPrint.Total})
	}
}

//GetCarbonPrintForItem ... With an injected data access repo, get the carbon print value for a given item
func (fc FoodCarbonService) GetCarbonPrintForItem(ctx *gin.Context) {
	itemName := ctx.Param("itemName")

	itemCarbonPrint, err := fc.repo.GetCarbonPrintForItem(itemName)
	if err != nil {
		ctx.JSON(500, ctx.Error(err))
	} else {
		ctx.JSON(200, responses.ItemCarbonPrint{Name: itemCarbonPrint.Name, CarbonPrint: itemCarbonPrint.CarbonPrint})
	}
}

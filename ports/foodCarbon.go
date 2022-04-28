package ports

import (
	app "github.com/Scileos/food-carbonprint-service/app/foodCarbon"
	"github.com/gin-gonic/gin"
)

//FoodCarbonController ... FoodCarbon Controller type
type FoodCarbonController struct {
	FoodCarbonService *app.FoodCarbonService
}

//Init ... Initialize user routes
func (c *FoodCarbonController) Init(router *gin.Engine) {
	router.POST("/api/carbonPrint/basket", c.FoodCarbonService.GetCarbonPrintForBasket)
	router.GET("/api/carbonPrint/item/:itemName", c.FoodCarbonService.GetCarbonPrintForItem)
}

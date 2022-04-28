package main

import (
	"fmt"
	"log"
	"net/http"

	app "github.com/Scileos/food-carbonprint-service/app/foodCarbon"
	"github.com/Scileos/food-carbonprint-service/config"
	"github.com/Scileos/food-carbonprint-service/ports"

	"github.com/Scileos/food-carbonprint-service/adapters"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))

	foodCarbonConfig := config.NewFoodCarbonConfig()
	foodCarbonRepo := adapters.NewFoodCarbonRepository(foodCarbonConfig)
	foodCarbonService := app.NewFoodCarbonService(foodCarbonRepo)

	//init controllers
	InitControllers(router, foodCarbonService)

	port := "8888"

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

// InitControllers ... Initialise avaiable controllers
func InitControllers(router *gin.Engine, foodCarbonService *app.FoodCarbonService) {
	foodCarbonController := ports.FoodCarbonController{FoodCarbonService: foodCarbonService}
	foodCarbonController.Init(router)
}

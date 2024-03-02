package main

import (
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/services"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/infrastucture/persistence/dynamoDB"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/interface/api/rest"

	"github.com/gin-gonic/gin"
)

func main() {

	dbClient, err := dynamoDB.CreateDynamoDBClient()

	if err != nil {
		fmt.Printf("Error initializing DynamoDB client: %v\n", err)
		return
	}

	gunplaRepo := dynamoDB.CreateGunplaRepository(dbClient.Client)
	toolRepo := dynamoDB.CreateToolRepository(dbClient.Client)
	orderRepo := dynamoDB.CreateOrderRepository(dbClient.Client)

	gunplaService := services.CreateGunplaService(gunplaRepo)
	toolService := services.CreateToolService(toolRepo)
	orderService := services.CreateOrderService(orderRepo)
	// orderService := services.CreateOrderService(orderRepo, gunplaRepo, toolRepo)

	gunplaController := rest.CreateGunplaController(gunplaService)
	toolController := rest.CreateToolController(toolService)
	orderController := rest.CreateOrderController(orderService)

	router := gin.Default()
	gunplaController.SetupRoutes(router)
	toolController.SetupRoutes(router)
	orderController.SetupRoutes(router)
	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

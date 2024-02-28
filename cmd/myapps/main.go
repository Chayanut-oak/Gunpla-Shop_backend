package main

import (
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/services"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/infrastucture/db/dynamoDB"
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
	// orderRepo := dynamoDB.CreateOrderRepository(dbClient.Client)

	gunplaService := services.CreateGunplaService(gunplaRepo)

	gunplaController := rest.CreateGunplaController(gunplaService)

	router := gin.Default()

	router.GET("/getAllGunpla", gunplaController.GetAllGunplasHandler)
	router.POST("/addGunpla", gunplaController.AddGunplaHHandler)
	router.PUT("/updateGunpla", gunplaController.UpdateGunplaHandler)
	router.DELETE("/deleteGunpla", gunplaController.DeleteGunplaHandler)
	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

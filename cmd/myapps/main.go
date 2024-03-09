package main

import (
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/services"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/services/auth"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/infrastucture/persistence/dynamoDB"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/infrastucture/persistence/s3"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/interface/api/rest"
	"github.com/gin-contrib/cors"
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
	userRepo := dynamoDB.CreateUserRepository(dbClient.Client)
	forumRepo := dynamoDB.CreateForumRepository(dbClient.Client)

	gunplaService := services.CreateGunplaService(gunplaRepo)
	toolService := services.CreateToolService(toolRepo)
	orderService := services.CreateOrderService(orderRepo)
	userService := services.CreateUserService(userRepo, auth.AuthService{})
	forumService := services.CreateForumService(forumRepo)
	// orderService := services.CreateOrderService(orderRepo, gunplaRepo, toolRepo)

	gunplaController := rest.CreateGunplaController(gunplaService)
	toolController := rest.CreateToolController(toolService)
	orderController := rest.CreateOrderController(orderService)
	userController := rest.CreateUserController(userService)
	forumController := rest.CreateForumController(forumService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	router.Use(cors.New(config))

	gunplaController.SetupRoutes(router)
	toolController.SetupRoutes(router)
	orderController.SetupRoutes(router)
	userController.SetupRoutes(router)
	forumController.SetupRoutes(router)
	router.GET("/api", s3.Pong)
	router.POST("/api/s3/upload-image", s3.S3uploader)
	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

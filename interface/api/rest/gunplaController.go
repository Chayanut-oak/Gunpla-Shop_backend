package rest

import (
	"fmt"
	"net/http"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/interfaces"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/services/auth"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type GunplaController struct {
	gunplaService interfaces.GunplaService
}

func CreateGunplaController(gunplaService interfaces.GunplaService) *GunplaController {
	return &GunplaController{
		gunplaService: gunplaService,
	}
}
func (gc *GunplaController) SetupRoutes(router *gin.Engine) {
	gunplaGroup := router.Group("/gunpla")
	{
		gunplaGroup.GET("", gc.GetAllGunplasHandler)
		gunplaGroup.Use(middleware.AuthMiddleware(&auth.AuthService{}))
		gunplaGroup.POST("/addGunpla", gc.AddGunplaHandler)
		gunplaGroup.PUT("/updateGunpla", gc.UpdateGunplaHandler)
		gunplaGroup.DELETE("/deleteGunpla/:productId", gc.DeleteGunplaHandler)
	}
}

func (controller *GunplaController) GetAllGunplasHandler(c *gin.Context) {
	gunplas, err := controller.gunplaService.GetAllGunplas()
	fmt.Println("from controller: ", gunplas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gunplas"})
		return
	}

	c.JSON(http.StatusOK, gunplas)
}

func (controller *GunplaController) AddGunplaHandler(c *gin.Context) {
	var gunpla restModel.GunplaRestModel
	role, exists := c.Get("role")
	fmt.Print(role)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	// Bind the JSON payload from the request body to the Gunpla struct
	if err := c.BindJSON(&gunpla); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Call the AddGunpla method of the gunplaService
	res, err := controller.gunplaService.AddGunpla(gunpla)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Gunpla item"})
		return
	}

	// Respond with the added Gunpla and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Gunpla item added successfully", "gunpla": res})
}
func (controller *GunplaController) UpdateGunplaHandler(c *gin.Context) {
	var gunpla entity.Gunpla
	role, exists := c.Get("role")
	fmt.Print(role)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	// Bind the JSON payload from the request body to the Gunpla struct
	if err := c.BindJSON(&gunpla); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	res, err := controller.gunplaService.UpdateGunpla(gunpla)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Gunpla item"})
		return
	}

	// Respond with the added Gunpla and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Gunpla item update successfully", "gunpla": res})
}

func (controller *GunplaController) DeleteGunplaHandler(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}
	if role != "admin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
		return
	}
	ProductId := c.Param("productId")
	err := controller.gunplaService.DeleteGunpla(ProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Gunpla item"})
		return
	}

	// Respond with the added Gunpla and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Gunpla item delete successfully"})
}

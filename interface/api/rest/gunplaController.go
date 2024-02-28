package rest

import (
	"fmt"
	"net/http"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/interfaces"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
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

func (controller *GunplaController) GetAllGunplasHandler(c *gin.Context) {
	gunplas, err := controller.gunplaService.GetAllGunplas()
	fmt.Println("from controller: ", gunplas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gunplas"})
		return
	}

	c.JSON(http.StatusOK, gunplas)
}

func (controller *GunplaController) AddGunplaHHandler(c *gin.Context) {
	var gunpla entity.NewGunpla

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
	GunplaId := c.Param("gunplaId")
	err := controller.gunplaService.DeleteGunpla(GunplaId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Gunpla item"})
		return
	}

	// Respond with the added Gunpla and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Gunpla item delete successfully"})
}
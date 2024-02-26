package rest

import (
	"fmt"
	"net/http"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/interfaces"
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

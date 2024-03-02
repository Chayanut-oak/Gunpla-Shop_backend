package rest

import (
	"fmt"
	"net/http"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/interfaces"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/gin-gonic/gin"
)

type ToolController struct {
	toolService interfaces.ToolService
}

func CreateToolController(toolService interfaces.ToolService) *ToolController {
	return &ToolController{
		toolService: toolService,
	}
}
func (gc *ToolController) SetupRoutes(router *gin.Engine) {
	toolGroup := router.Group("/tool")
	{
		toolGroup.GET("", gc.GetAllToolsHandler)
		toolGroup.POST("/addTool", gc.AddToolHandler)
		toolGroup.PUT("/updateTool", gc.UpdateToolHandler)
		toolGroup.DELETE("/deleteTool/:productId", gc.DeleteToolHandler)
	}
}

func (controller *ToolController) GetAllToolsHandler(c *gin.Context) {
	tools, err := controller.toolService.GetAllTools()
	fmt.Println("from controller: ", tools)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tools"})
		return
	}

	c.JSON(http.StatusOK, tools)
}

func (controller *ToolController) AddToolHandler(c *gin.Context) {
	var tool restModel.ToolRestModal

	// Bind the JSON payload from the request body to the Tool struct
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Call the AddTool method of the toolService
	res, err := controller.toolService.AddTool(tool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Tool item"})
		return
	}

	// Respond with the added Tool and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Tool item added successfully", "tool": res})
}
func (controller *ToolController) UpdateToolHandler(c *gin.Context) {
	var tool entity.Tool

	// Bind the JSON payload from the request body to the Tool struct
	if err := c.BindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	res, err := controller.toolService.UpdateTool(tool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Tool item"})
		return
	}

	// Respond with the added Tool and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Tool item update successfully", "tool": res})
}

func (controller *ToolController) DeleteToolHandler(c *gin.Context) {
	ProductId := c.Param("productId")
	err := controller.toolService.DeleteTool(ProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Tool item"})
		return
	}

	// Respond with the added Tool and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Tool item delete successfully"})
}

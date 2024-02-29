package rest

import (
	"fmt"
	"net/http"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/application/interfaces"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService interfaces.OrderService
}

func CreateOrderController(orderService interfaces.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}
func (oc *OrderController) SetupRoutes(router *gin.Engine) {
	orderaGroup := router.Group("/order")
	{
		orderaGroup.GET("", oc.GetAllOrdersHandler)
		orderaGroup.POST("/addOrder", oc.AddOrderHandler)
		orderaGroup.PUT("/updateOrder", oc.UpdateOrderHandler)
		orderaGroup.DELETE("/deleteOrder/:OrderId", oc.DeleteOrderHandler)
	}
}

func (controller *OrderController) GetAllOrdersHandler(c *gin.Context) {
	orders, err := controller.orderService.GetAllOrders()
	fmt.Println("from controller: ", orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (controller *OrderController) AddOrderHandler(c *gin.Context) {
	var order restModel.OrderRestModal
	// Bind the JSON payload from the request body to the order struct
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Call the Addorder method of the orderService
	res, err := controller.orderService.AddOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order item"})
		return
	}

	// Respond with the added order and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "order item added successfully", "order": res})
}
func (controller *OrderController) UpdateOrderHandler(c *gin.Context) {
	var order entity.Order

	// Bind the JSON payload from the request body to the order struct
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	res, err := controller.orderService.UpdateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Gunpla item"})
		return
	}

	// Respond with the added Gunpla and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Gunpla item update successfully", "gunpla": res})
}

func (controller *OrderController) DeleteOrderHandler(c *gin.Context) {
	OrderId := c.Param("orderId")
	err := controller.orderService.DeleteOrder(OrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Gunpla item"})
		return
	}

	// Respond with the added Gunpla and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Gunpla item delete successfully"})
}
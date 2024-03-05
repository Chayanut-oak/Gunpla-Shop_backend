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

type UserController struct {
	userService interfaces.UserService
}
type AuthResponse struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}

func CreateUserController(userService interfaces.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (gc *UserController) SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("", middleware.AuthMiddleware(&auth.AuthService{}), gc.GetUserHandler)
		userGroup.POST("/newUser", gc.NewUserHandler)
		userGroup.POST("/authentication", gc.Authentication)
		userGroup.GET("/allUser", middleware.AuthMiddleware(&auth.AuthService{}), gc.GetAllUserHandler)
	}
}

func (controller *UserController) NewUserHandler(c *gin.Context) {
	var newUser restModel.UserRestModel
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(newUser.Address)
	res, err := controller.userService.NewUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new user"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (controller *UserController) Authentication(c *gin.Context) {
	var authRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := controller.userService.AuthenticateUser(authRequest.Email, authRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	response := AuthResponse{
		Token: token,
		User:  user,
	}
	c.JSON(http.StatusOK, response)

}
func (controller *UserController) GetUserHandler(c *gin.Context) {
	email, exists := c.Get("email")
	fmt.Print(email)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	user, err := controller.userService.GetUser(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Customer"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func (controller *UserController) GetAllUserHandler(c *gin.Context) {

	allUser, err := controller.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Customer"})
		return
	}

	c.JSON(http.StatusOK, allUser)
}

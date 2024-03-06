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

type ForumController struct {
	forumService interfaces.ForumService
}

func CreateForumController(forumService interfaces.ForumService) *ForumController {
	return &ForumController{
		forumService: forumService,
	}
}
func (gc *ForumController) SetupRoutes(router *gin.Engine) {
	forumGroup := router.Group("/forum")
	{
		forumGroup.GET("", gc.GetAllForumsHandler)
		forumGroup.GET("/comment/:forumId", gc.GetAllCommentsInForumHandler)
		forumGroup.Use(middleware.AuthMiddleware(&auth.AuthService{}))
		forumGroup.POST("/addForum", gc.AddForumHandler)
		forumGroup.PUT("/updateForum", gc.UpdateForumHandler)
		forumGroup.DELETE("/deleteForum/:forumId", gc.DeleteForumHandler)
		forumGroup.POST("/addComment", gc.AddCommentHandler)
		forumGroup.PUT("/updateComment", gc.UpdateCommentHandler)
		forumGroup.DELETE("/deleteComment/:commentId", gc.DeleteCommentHandler)
	}
}

func (controller *ForumController) GetAllForumsHandler(c *gin.Context) {
	forums, err := controller.forumService.GetAllForums()
	fmt.Println("from controller: ", forums)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch forums"})
		return
	}

	c.JSON(http.StatusOK, forums)
}

func (controller *ForumController) AddForumHandler(c *gin.Context) {
	var forum restModel.ForumRestModel
	// role, exists := c.Get("role")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
	// 	return
	// }
	// if role != "admin" {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
	// 	return
	// }
	// Bind the JSON payload from the request body to the Forum struct
	if err := c.BindJSON(&forum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Call the AddForum method of the forumService
	res, err := controller.forumService.AddForum(forum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Forum item"})
		return
	}

	// Respond with the added Forum and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Forum item added successfully", "forum": res})
}
func (controller *ForumController) UpdateForumHandler(c *gin.Context) {
	var forum entity.Forum

	// Bind the JSON payload from the request body to the Forum struct
	if err := c.BindJSON(&forum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	res, err := controller.forumService.UpdateForum(forum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Forum item"})
		return
	}

	// Respond with the added Forum and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Forum item update successfully", "forum": res})
}

func (controller *ForumController) DeleteForumHandler(c *gin.Context) {
	// role, exists := c.Get("role")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
	// 	return
	// }

	ForumId := c.Param("forumId")
	err := controller.forumService.DeleteForum(ForumId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Forum item"})
		return
	}

	// Respond with the added Forum and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Forum item delete successfully"})
}

func (controller *ForumController) GetAllCommentsInForumHandler(c *gin.Context) {
	ForumId := c.Param("forumId")
	forum, err := controller.forumService.GetForum(ForumId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch forum"})
		return
	}
	fmt.Println("forum", forum)
	comments, err := controller.forumService.GetAllCommentsInForum(ForumId)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"forum": forum, "comments": comments})

}

func (controller *ForumController) AddCommentHandler(c *gin.Context) {
	var comment restModel.CommentRestModel
	// role, exists := c.Get("role")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
	// 	return
	// }
	// if role != "admin" {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have permission"})
	// 	return
	// }
	// Bind the JSON payload from the request body to the Forum struct
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Call the AddForum method of the forumService
	res, err := controller.forumService.AddComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Forum item"})
		return
	}

	// Respond with the added Forum and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Comment item added successfully", "forum": res})
}
func (controller *ForumController) UpdateCommentHandler(c *gin.Context) {
	var comment entity.Comment

	// Bind the JSON payload from the request body to the Forum struct
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	res, err := controller.forumService.UpdateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Comment item"})
		return
	}

	// Respond with the added Forum and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Comment item update successfully", "Comment": res})
}

func (controller *ForumController) DeleteCommentHandler(c *gin.Context) {
	// role, exists := c.Get("role")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
	// 	return
	// }

	CommentId := c.Param("commentId")
	err := controller.forumService.DeleteComment(CommentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Comment item"})
		return
	}

	// Respond with the added Forum and a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Comment item delete successfully"})
}

package routes

import (
	"net/http"

	"github.com/eric-batista/neowire-api-task-manager-poc/application/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTaskRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/tasks", func(c *gin.Context) {
		if response, err := usecases.CreateNewTask(c, db); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, response)
		}
	})
	r.GET("/tasks", func(c *gin.Context) {})
	r.GET("/tasks/:id", func(c *gin.Context) {})
	r.PATCH("/tasks/:id", func(c *gin.Context) {})
	r.DELETE("/tasks/:id", func(c *gin.Context) {})
}

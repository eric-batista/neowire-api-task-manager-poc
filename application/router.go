package application

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/application/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApplicationRouter(r *gin.RouterGroup, db *gorm.DB) {
	routes.RegisterTaskRoutes(r, db)
}

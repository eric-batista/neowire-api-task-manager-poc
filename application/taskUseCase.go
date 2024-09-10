package application

import (
	"net/http"

	"github.com/eric-batista/neowire-api-task-manager-poc/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Função para registrar as rotas de tarefas
func RegisterTaskRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/tasks", func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&task)
		c.JSON(http.StatusOK, task)
	})

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []models.Task
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		id := c.Param("id")
		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		id := c.Param("id")
		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&task)
		c.JSON(http.StatusOK, task)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		id := c.Param("id")
		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
			return
		}

		db.Delete(&task)
		c.JSON(http.StatusOK, gin.H{"message": "Tarefa excluída com sucesso"})
	})
}

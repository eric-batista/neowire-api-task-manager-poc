package usecases

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	"github.com/eric-batista/neowire-api-task-manager-poc/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewTask(c *gin.Context, db *gorm.DB) (*models.Task, error) {
	var payload models.TaskPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, err
	}

	task, err := models.NewTask(payload.Title, payload.Description, payload.Status)
	if err != nil {
		return nil, err
	}

	taskRepository := repositories.TasksRepositortyDB{DB: db}
	if err := taskRepository.AddTask(task); err != nil {
		return nil, err
	}

	return task, nil
}

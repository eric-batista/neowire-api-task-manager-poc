package usecases

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	"github.com/eric-batista/neowire-api-task-manager-poc/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateTask(c *gin.Context, db *gorm.DB, taskID string) (*models.Task, error) {
	var payload models.TaskPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, err
	}

	taskRepository := repositories.TasksRepositortyDB{DB: db}
	task, err := taskRepository.RetrieveTask(taskID)
	if err != nil {
		return nil, err
	}

	updatedTask, err := models.UpdateTask(task, payload.Title, payload.Description, payload.Status, payload.ClientID)
	if err != nil {
		return nil, err
	}

	if newTask, err := taskRepository.UpdateTask(updatedTask, taskID); err != nil {
		return nil, err
	} else {
		return newTask, nil
	}
}

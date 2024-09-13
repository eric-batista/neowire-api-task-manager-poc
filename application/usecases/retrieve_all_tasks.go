package usecases

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	"github.com/eric-batista/neowire-api-task-manager-poc/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RetrieveAllTasks(c *gin.Context, db *gorm.DB) (*[]models.Task, error) {
	var payload models.TaskListPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		return nil, err
	}

	taskRepository := repositories.TasksRepositortyDB{DB: db}
	if tasks, err := taskRepository.RetrieveAllTasks(payload.ClientID); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}

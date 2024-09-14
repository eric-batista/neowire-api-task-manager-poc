package usecases

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	"github.com/eric-batista/neowire-api-task-manager-poc/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RetrieveTask(c *gin.Context, db *gorm.DB, taskID string) (*models.Task, error) {
	taskRepository := repositories.TasksRepositortyDB{DB: db}
	if task, err := taskRepository.RetrieveTask(taskID); err != nil {
		return nil, err
	} else {
		return task, nil
	}
}

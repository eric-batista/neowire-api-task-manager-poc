package usecases

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTask(c *gin.Context, db *gorm.DB, taskID string) error {
	taskRepository := repositories.TasksRepositortyDB{DB: db}
	if err := taskRepository.DeleteTask(taskID); err != nil {
		return err
	}

	return nil
}

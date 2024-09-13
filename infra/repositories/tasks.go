package repositories

import (
	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	"gorm.io/gorm"
)

type TasksRepositortyDB struct {
	DB *gorm.DB
}

func (r TasksRepositortyDB) AddTask(task *models.Task) error {
	if err := r.DB.Create(task); err != nil {
		return err.Error
	}
	return nil
}

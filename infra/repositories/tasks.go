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

func (r TasksRepositortyDB) RetrieveAllTasks(clientID string) (*[]models.Task, error) {
	var result []models.Task
	if err := r.DB.Table("tasks").Where("client_id = ?", clientID).Scan(&result); err.Error != nil {
		return nil, err.Error
	}
	return &result, nil
}

func (r TasksRepositortyDB) RetrieveTask(taskID string) (*models.Task, error) {
	var result models.Task
	if err := r.DB.Table("tasks").Where("id = ?", taskID).Scan(&result); err.Error != nil {
		return nil, err.Error
	}
	return &result, nil
}

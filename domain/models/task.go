package models

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	Base        `validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

func (task *Task) isValid() error {
	validator := validator.New()
	err := validator.Struct(task)
	if err != nil {
		return err
	}
	return nil
}

func NewTask(title, description, status, clientId string) (*Task, error) {
	task := Task{
		Title:       title,
		Description: description,
		Status:      status,
	}

	task.ID = uuid.NewV4().String()
	task.ClientID = clientId
	task.CreatedAt = time.Now()

	err := task.isValid()
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *Task, title, description, status, clientId string) (*Task, error) {
	newTask := Task{
		Title:       title,
		Description: description,
		Status:      status,
	}

	newTask.ID = task.ID
	newTask.ClientID = task.ClientID
	newTask.CreatedAt = task.CreatedAt
	newTask.UpdatedAt = time.Now()

	err := newTask.isValid()
	if err != nil {
		return nil, err
	}

	return &newTask, err
}

type TaskPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ClientID    string `json:"clientId"`
}

type TaskListPayload struct {
	ClientID string `json:"clientId"`
}

type RetrieveTask struct {
	TaskID string `json:"taskId"`
}

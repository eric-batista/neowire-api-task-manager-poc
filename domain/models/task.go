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

func NewTask(title, description, status string) (*Task, error) {
	task := Task{
		Title:       title,
		Description: description,
		Status:      status,
	}

	task.ID = uuid.NewV4().String()
	task.CreatedAt = time.Now()

	err := task.isValid()
	if err != nil {
		return nil, err
	}
	return &task, nil
}

type TaskPayload struct {
	Title       string
	Description string
	Status      string
}

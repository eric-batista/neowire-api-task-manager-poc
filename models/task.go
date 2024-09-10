package models

import "gorm.io/gorm"

// Estrutura da Task (Tarefa)
type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

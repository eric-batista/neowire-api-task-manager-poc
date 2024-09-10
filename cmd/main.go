package main

import (
	"net/http"

	"github.com/eric-batista/neowire-api-task-manager-poc/application"
	"github.com/eric-batista/neowire-api-task-manager-poc/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Inicializa a conexão com o banco de dados (SQLite para simplificação)
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Faz a migração do modelo de Task
	db.AutoMigrate(&models.Task{})

	// Inicializa o router Gin
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Registra os endpoints da API
	application.RegisterTaskRoutes(r, db)

	// Inicia o servidor
	r.Run(":8080") // Executa na porta 8080
}

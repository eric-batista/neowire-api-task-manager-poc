package main

import (
	"net/http"

	"github.com/eric-batista/neowire-api-task-manager-poc/application"
	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	"github.com/eric-batista/neowire-api-task-manager-poc/infra/core"
	infra "github.com/eric-batista/neowire-api-task-manager-poc/infra/database"
	"github.com/gin-gonic/gin"
)

func main() {
	baseConfig := core.SetEnvs()

	dbConfig := infra.CreateDBConfig()
	db, err := infra.ConnectDB(dbConfig)
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		panic("Falha ao rodar migrations")
	}

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

	router := r.Group(baseConfig.BasePath)
	router_v1 := router.Group("/v1")
	application.ApplicationRouter(router_v1, db)
	r.Run(":8080") // Executa na porta 8080
}

package database

import (
	"fmt"
	"log"

	"github.com/eric-batista/neowire-api-task-manager-poc/infra/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func CreateDBConfig() *DBConfig {
	return &DBConfig{
		Host:     core.GetEnv("DB_HOST", "localhost"),
		Port:     core.GetEnvAsInt("DB_PORT", 5432),
		User:     core.GetEnv("DB_USER", "postgres"),
		Password: core.GetEnv("DB_PASSWOR", "postgres"),
		DBName:   core.GetEnv("DB_NAME", "tasks-manager"),
		SSLMode:  core.GetEnv("DB_SSLMODE", "disable"),
	}
}

func ConnectDB(cfg *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Conectado ao banco de dados com sucesso!")
	return db, nil
}

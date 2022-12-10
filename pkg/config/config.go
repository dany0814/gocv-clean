package config

import (
	"fmt"
	"gocv-clean/pkg/models"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	// Database config
	DbUser    string `default:"postgres" required:"true"`
	DbPass    string `default:"postgres" required:"true"`
	DbHost    string `default:"localhost" required:"true"`
	DbPort    uint   `default:"5432" required:"true"`
	DbName    string `default:"postgres" required:"true"`
	DbTimeout time.Duration
	// Server config
	Host string `default:"localhost" required:"true"`
	Port uint   `default:"3000" required:"true"`
	// ShutdownTimeout time.Duration
}

var Env config

func LoadConfig() error {
	err := envconfig.Process("DANY", &Env)
	if err != nil {
		return err
	}
	return nil
}

func NewDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", Env.DbHost, Env.DbUser, Env.DbPass, Env.DbPass, Env.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("couldn't connect to postgresql: %v", err)
	}

	db.AutoMigrate(&models.Book{})

	return db, nil
}

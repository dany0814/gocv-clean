package bootstrap

import (
	"gocv-clean/internal/platform/server"
	"gocv-clean/pkg/config"
	"log"
)

const (
	host = "localhost"
	port = 3000
)

func Run() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	config.NewDb()

	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}
	srv := server.New(host, port)
	return srv.Run()
}

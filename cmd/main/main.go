package main

import (
	"log"
	"os"

	restapi "github.com/dCastillo727/go-architecture/internal/driving/rest_api"
	"github.com/dCastillo727/go-architecture/internal/registry"
	"github.com/joho/godotenv"
)

func main() {
	cfg := loadConfig()
	app, _ := composeApp(cfg)

	defer app.Close()

	router := restapi.SetupRouter(app.Services)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func loadConfig() *registry.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Environment file not found, using system values")
	}

	return &registry.Config{
		DdbbHost:     os.Getenv("DDBB_HOST"),
		DdbbUser:     os.Getenv("DDBB_USER"),
		DdbbPassword: os.Getenv("DDBB_PASSWORD"),
		DdbbName:     os.Getenv("DDBB_NAME"),
	}
}

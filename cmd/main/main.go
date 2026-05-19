package main

import (
	restapi "github.com/dCastillo727/go-architecture/internal/driving/rest_api"
)

func main() {
	app := composeApp()

	router := restapi.SetupRouter(app.Services)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

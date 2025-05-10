package main

import (
	"log"
	"os"
	"quiz-app/app"
	"quiz-app/infrastructure/config"
)

func main() {
	// Bootstrap configuration
	if err := config.Load(); err != nil {
		log.Println("failed to load configurations: ", err)
		os.Exit(1)
	}

	// Start the application
	svc, err := app.NewService()
	if err != nil {
		log.Println("failed to start login-service [POLARIS]: ", err)
		os.Exit(1)
	}

	// Start the service
	svc.Run()
}

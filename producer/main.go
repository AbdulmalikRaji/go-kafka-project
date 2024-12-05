package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdulmalikraji/go-kafka-project/producer/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	client := config.NewConnection()
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Error closing Kafka connection: %v\n", err)
		}
	}()

	config.InitializeRoutes(app, client)

	// Start the server in a goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Call gracefulShutdown to handle cleanup
	gracefulShutdown(app, client)
}

func gracefulShutdown(app *fiber.App, client *config.Client) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// Attempt to shut down the Fiber app gracefully
	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	}

	// Close the Kafka connection
	if err := client.Close(); err != nil {
		log.Printf("Error closing Kafka connection: %v\n", err)
	}

	log.Println("Server shutdown complete.")
}

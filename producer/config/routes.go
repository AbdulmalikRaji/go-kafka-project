package config

import (
	handler "github.com/abdulmalikraji/go-kafka-project/producer/Handler"
	"github.com/abdulmalikraji/go-kafka-project/producer/services"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App, client *Client) {

	//Initialize Service
	commentService := services.NewCommentService(client.SamaraConnection)

	//Initialize Controller
	commentHandler := handler.NewCommentController(commentService)

	// Initialize routes
	api := app.Group("/api/v1")
	api.Post("/comment", commentHandler.CreateComment)

}

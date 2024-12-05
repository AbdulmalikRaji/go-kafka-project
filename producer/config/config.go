package config

import (
	"sync"

	"github.com/IBM/sarama"
	"github.com/abdulmalikraji/go-kafka-project/producer/services"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

type Client struct {
	SamaraConnection sarama.SyncProducer
}

var client *Client

func NewConnection() *Client {
	once.Do(func() {
		client = &Client{
			SamaraConnection: ConnectProducer(),
		}
	})

	return client
}

func InitializeRoutes(app *fiber.App, client *Client) {

	commentService := services.NewCommentService(client.SamaraConnection)

	// Initialize routes 
	api := app.Group("/api/v1")
	api.Post("/comment", commentService.CreateComment)

}

func (c *Client) Close() error {
	return c.SamaraConnection.Close()
}
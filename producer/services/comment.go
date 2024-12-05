package services

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/abdulmalikraji/go-kafka-project/producer/dto"
	"github.com/gofiber/fiber/v2"
)

type CommentService interface {
	PushCommentToQueue(topic string, message []byte) error
	CreateComment(c *fiber.Ctx) error
}

type commentService struct {
	producer sarama.SyncProducer
}

func NewCommentService(producer sarama.SyncProducer) CommentService {
	return &commentService{producer: producer}
}

func (s commentService) PushCommentToQueue(topic string, message []byte) error {

	// Convert the comment to a JSON string in samara format
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send the message
	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Message sent to topic %s partition %d at offset %d\n", topic, partition, offset)

	return nil
}

func (s commentService) CreateComment(c *fiber.Ctx) error {
	var comment dto.Comment
	err := c.BodyParser(&comment)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	commentInBytes, err := json.Marshal(comment)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	s.PushCommentToQueue("comment", commentInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "comment added successfully",
		"data":    comment,
	})

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return err
}

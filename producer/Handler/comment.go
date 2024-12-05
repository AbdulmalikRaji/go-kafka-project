package handler

import (
	"encoding/json"
	"log"

	"github.com/abdulmalikraji/go-kafka-project/producer/dto"
	"github.com/abdulmalikraji/go-kafka-project/producer/services"
	"github.com/gofiber/fiber/v2"
)

type CommentController interface {
	CreateComment(c *fiber.Ctx) error
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return commentController{
		service: service,
	}
}

func (s commentController) CreateComment(c *fiber.Ctx) error {
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

	s.service.PushCommentToQueue("comment", commentInBytes)

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

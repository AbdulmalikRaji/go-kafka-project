package services

import (
	"log"

	"github.com/IBM/sarama"
)

type CommentService interface {
	PushCommentToQueue(topic string, message []byte) error
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

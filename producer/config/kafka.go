package config

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

func ConnectProducer() sarama.SyncProducer {
	//setup samara client connection
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	brokerUrl := []string{os.Getenv("BROKER_URL")} 

	producer, err := sarama.NewSyncProducer(brokerUrl, config)
	if err != nil {
		log.Println("Failed to connect to Kafka:", err)
		return nil
	}

	return producer
}

package config

import (
    "fmt"
    "log"
    "os"

    "github.com/IBM/sarama"
    "github.com/joho/godotenv" 
)

func ConnectProducer() sarama.SyncProducer {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file:", err)
    }

    // Setup sarama client connection
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true

    brokerUrl := os.Getenv("BROKER_URL") // Load BROKER_URL from .env
    fmt.Println("Connecting to", brokerUrl)

    producer, err := sarama.NewSyncProducer([]string{brokerUrl}, config)
    if err != nil {
        log.Println("Failed to connect to Kafka:", err)
        return nil
    }

    return producer
}

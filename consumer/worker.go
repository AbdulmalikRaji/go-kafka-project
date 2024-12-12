package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {

	topic := os.Getenv("TOPIC")
	worker, err := connectConsumer([]string{os.Getenv("BROKER_URL")})
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0

	donech := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)

			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Recieved %d messages, | Topic: %s | Message: %s \n", msgCount, string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interruped")
				donech <- struct{}{}
				return
			}
		}

	}()

	<-donech
	fmt.Println("Processed", msgCount, " messages")
	if err := worker.Close(); err != nil {
		log.Fatalf("Error closing Kafka connection: %v\n", err)
	}

}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
    config.Consumer.Return.Errors = true

    client, err := sarama.NewConsumer(brokers, config)
    if err!= nil {
        return nil, err
    }

    return client, nil
}
# Go Kafka Event-Based Architecture

This project is for learning event-based architecture using Go and Kafka. It implements a simple producer-consumer (pub-sub) model where comments with a are sent from a producer to a consumer.

The focus is on understanding how to integrate Kafka into Go applications to handle asynchronous messaging.

## Setup

1. Install Docker and Docker Compose

2. Run `docker-compose up -d` in the root directory. This will start Apache Kafka and Zookeeper

3. Install dependencies

4. Run `go run producer/main.go` to start the producer which is a REST API listening on port 3000

5. Run `go run worker/worker.go` to start the consumer

### Test it out

Send a POST request to `localhost:3000`

This will produce a message in Apache Kafka and the consumer will process it.

```bash
curl --location --request POST '0.0.0.0:3000/api/v1/comment' \
--header 'Content-Type: application/json' \
--data-raw '{ "text":"message 1" }'

curl --location --request POST '0.0.0.0:3000/api/v1/comment' \
--header 'Content-Type: application/json' \
--data-raw '{ "text":"message 2" }'
```

package config

import (
	"sync"

	"github.com/IBM/sarama"
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

func (c *Client) Close() error {
	return c.SamaraConnection.Close()
}

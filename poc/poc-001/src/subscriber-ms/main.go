package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
)

const (
	topic          = "message-log"
	broker1Address = "localhost:9092"
)

func main() {
	//app := fiber.New()

	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Subscribe to subject

	nc.Subscribe("updates", func(msg *nats.Msg) {
		// Handle the message
		log.Printf("NATS: Listening on [updates]")
		log.Printf("Received message '%s'", string(msg.Data))
	})

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "my-topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}()

	// Create a RabbitMQ connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Create a consumer
	consumer, err := ch.Consume("my-queue", "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for delivery := range consumer {
		// Get the message from the delivery
		message := delivery.Body

		// Process the message
		fmt.Println(string(message))

		// Acknowledge the message
		delivery.Ack(false)

	}

	// Keep the connection alive

}

func consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}

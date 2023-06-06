package provider

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"

	"routerms/model"
)

type rabbitMQProvider struct {
	AmqpConn *amqp.Connection
	Queue    amqp.Queue
}

func NewRabbitMQProvider() *rabbitMQProvider {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Create a queue
	q, err := ch.QueueDeclare("my-queue", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	return &rabbitMQProvider{
		AmqpConn: conn,
		Queue:    q,
	}
}

func (rs *rabbitMQProvider) Send(c *fiber.Ctx) error {
	m := new(model.Message)
	if err := c.BodyParser(m); err != nil {
		return err
	}

	ch, err := rs.AmqpConn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Send message to RabbitMQ
	ch.Publish("", rs.Queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(m.Message),
	})

	return c.JSON(fiber.Map{
		"message": "RabbitMQ: Message sent successfully",
	})
}

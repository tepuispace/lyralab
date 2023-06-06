package provider

import (
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"

	"routerms/model"
)

type kafkaProvider struct {
	KafkaSrv *kafka.Writer
}

func NewKafkaProvider() *kafkaProvider {
	// Kafka writer setup
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "kafka-topic",
		Balancer: &kafka.LeastBytes{},
	}
	return &kafkaProvider{
		KafkaSrv: w,
	}
}

func (ks *kafkaProvider) Send(c *fiber.Ctx) error {
	m := new(model.Message)
	if err := c.BodyParser(m); err != nil {
		return err
	}

	// Simple Publisher
	ks.KafkaSrv.WriteMessages(c.Context(), kafka.Message{
		Value: []byte(m.Message),
	})

	return c.JSON(fiber.Map{
		"message": "Message sent successfully",
	})
}

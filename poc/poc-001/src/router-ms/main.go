package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"routerms/provider"
)

func main() {
	app := fiber.New()
	// NATS
	ns := provider.NewNatsProvider()
	app.Post("/nats/send", ns.Send)

	// KAFKA
	ks := provider.NewKafkaProvider()
	app.Post("/kafka/send", ks.Send)

	// RABBIT
	rbs := provider.NewRabbitMQProvider()
	app.Post("/rabbitmq/send", rbs.Send)

	log.Fatal(app.Listen(":3000"))
}

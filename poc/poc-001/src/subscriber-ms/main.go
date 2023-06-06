package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	app := fiber.New()

	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Subscribe to subject

	nc.Subscribe("updates", func(msg *nats.Msg) {
		// Handle the message
		log.Printf("NATS: Listening on [updates]")
		log.Printf("Received message '%s'", string(msg.Data))
	})

	// Keep the connection alive
	log.Fatal(app.Listen(":3001"))
}

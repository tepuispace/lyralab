package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"routerms/provider"
)

func main() {
	app := fiber.New()
	//nc, _ := nats.Connect(nats.DefaultURL)
	ns := provider.NewNatsProvider()
	/*
		app.Post("/mats/send", func(c *fiber.Ctx) error {
			m := new(Message)
			if err := c.BodyParser(m); err != nil {
				return err
			}

			// Simple Publisher
			nc.Publish("updates", []byte(m.Message))

			return c.JSON(fiber.Map{
				"message": "Message sent successfully",
			})
		})
	*/
	app.Post("/mats/send", ns.Send)

	log.Fatal(app.Listen(":3000"))
}

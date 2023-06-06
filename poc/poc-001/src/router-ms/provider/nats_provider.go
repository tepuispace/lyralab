package provider

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"

	"routerms/model"
)

type natsProvider struct {
	NatConn *nats.Conn
}

func NewNatsProvider() *natsProvider {
	ns, _ := nats.Connect(nats.DefaultURL)
	return &natsProvider{
		NatConn: ns,
	}
}

func (ns *natsProvider) Send(c *fiber.Ctx) error {
	m := new(model.Message)
	if err := c.BodyParser(m); err != nil {
		return err
	}

	// Simple Publisher
	ns.NatConn.Publish("updates", []byte(m.Message))

	return c.JSON(fiber.Map{
		"message": "Message sent successfully",
	})
}

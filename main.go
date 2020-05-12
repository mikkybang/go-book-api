package main

import (
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(8000)
}

package main

import (
	"github.com/gofiber/fiber"
	"github.com/mikkybang/go-book-api/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(8000)
}

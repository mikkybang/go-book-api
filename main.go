package main

import (
	"github.com/gofiber/fiber"
	"github.com/mikkybang/go-book-api/book"
	"github.com/mikkybang/go-book-api/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:1d", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", helloWorld)

	app.Listen(8000)
}

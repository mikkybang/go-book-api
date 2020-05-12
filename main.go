package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/mikkybang/go-book-api/book"
	"github.com/mikkybang/go-book-api/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:1d", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection sucessfully opened")

}

func main() {
	app := fiber.New()
	initDatabase()

	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(8000)
}

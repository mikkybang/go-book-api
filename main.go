package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mikkybang/go-book-api/book"
	"github.com/mikkybang/go-book-api/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:1d", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("api/v1/book/:id", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection sucessfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(8000)
}

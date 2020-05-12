package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/mikkybang/go-book-api/database"
)

type Book struct {
	gorm.Model
	Title  string `json: "title"`
	Author string `json: "author"`
	Rating int    `json: "rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	db.Create(&book)
	c.JSON(book)
}

func DeleteBooks(c *fiber.Ctx) {
	c.Send("Delets a Book")
}

// func GetBooks(c *fiber.Ctx) {
// 	c.Send("All Books ")
// }

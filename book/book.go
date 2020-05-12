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
	c.Send("A Single Book ")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new Book")
}

func DeleteBooks(c *fiber.Ctx) {
	c.Send("Delets a Book")
}

// func GetBooks(c *fiber.Ctx) {
// 	c.Send("All Books ")
// }

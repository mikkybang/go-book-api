package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json: "title"`
	Author string `json: "author"`
	Rating int    `json: "rating"`
}

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books ")
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

package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books ")
}


func GetBook(c *fiber.Ctx) {
	c.Send("A Single Book ")
}


func NewBooks(c *fiber.Ctx) {
	c.Send("Adds a new Book")
}


func DeleteBooks(c *fiber.Ctx) {
	c.Send("Delets a Book")
}


// func GetBooks(c *fiber.Ctx) {
// 	c.Send("All Books ")
// }
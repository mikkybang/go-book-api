package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/mikkybang/go-book-api/database"
	"fmt"
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
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBooks(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book

	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	book := new(Book)
	
	fmt.Println(book)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Model(&book).Where("id = ?", id).Updates(book)
	c.JSON(book)
}

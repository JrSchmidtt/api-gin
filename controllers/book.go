package controllers

import (
	"strconv"

	"github.com/JrSchmidtt/api-gin/database"
	"github.com/JrSchmidtt/api-gin/models"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Not found" + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context){
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: "+ err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: "+ err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func GetAllBooks(c *gin.Context){
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot get all books: "+ err.Error(),
		})
		return
	}
	c.JSON(201, books)
}

func DeleteBook(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	
	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete book",
		})
		return
	}

	c.Status(204)
}
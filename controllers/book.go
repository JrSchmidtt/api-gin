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
}
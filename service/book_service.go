package service

import (
	"firstgo-p/database"
	"firstgo-p/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindBookById(c *gin.Context) {
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
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

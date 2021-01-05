package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/menezee/loyalty-go/src/db"
)

func peopleRoot(c *gin.Context) {
	persons := db.GetFirstNames()

	c.JSON(200, gin.H{
		"message": persons,
	})
}

func RegisterPeopleRoutes(router *gin.RouterGroup) {
	router.GET("/", peopleRoot)
}

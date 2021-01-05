package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) string {
	if max == 0 {
		max = 100
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max - min) + min
	return fmt.Sprintf("%03d", r)
}

func root(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})
	c.String(200, generateRandomNumber(0, 100))
}

func RegisterTokenRoutes(router *gin.RouterGroup) {
	router.GET("/", root)
}

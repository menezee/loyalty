package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/menezee/loyalty-go/src/db"
	"github.com/menezee/loyalty-go/src/handlers"
)

func main() {
	var err error
	_, err = db.New()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	handlers.RegisterTokenRoutes(r.Group("/generate-token"))
	handlers.RegisterPeopleRoutes(r.Group("/people"))

	r.Run()
}

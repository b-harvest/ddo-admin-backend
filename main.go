package main

import (
	"bharvest-vo/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	router.Start(r)
	r.Run("0.0.0.0:1317") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

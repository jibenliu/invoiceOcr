package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})
	r.Run(":8080")
}

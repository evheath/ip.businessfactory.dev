package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		addr := c.Request.RemoteAddr
		c.String(200, addr)
	})
	router.Run(":3000")
}

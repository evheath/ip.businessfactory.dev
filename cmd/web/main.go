package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		addr := c.Request.RemoteAddr
		c.String(http.StatusOK, addr)
	})
	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 12 * time.Second,
		IdleTimeout:  time.Second,
	}
	server.ListenAndServe()
}

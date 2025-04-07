package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		addr := c.Request.RemoteAddr
		if forwardAddr := c.Request.Header.Get("X-Forwarded-For"); forwardAddr != "" {
			// reverse proxy overwrites RemoteAddr but adds X-Forwarded-For
			parts := strings.Split(forwardAddr, ",")
			addr = strings.TrimSpace(parts[0])
		}
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

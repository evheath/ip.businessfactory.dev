package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
func main() {
	port := fmt.Sprintf(":%s", getEnv("PORT", "3000"))
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello foundry", port)
	})

	_ = http.ListenAndServe(port, nil)
}

package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	req := gin.Default()

	req.GET("/ping", PingHandler)

	err := req.Run()
	if err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}

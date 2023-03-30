package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yosikez/item-api/router"
)

func main() {
	r := gin.Default()

	router.Api(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	log.Println("Server started on port 8080")
}
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yehaozz/go-secure-api/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

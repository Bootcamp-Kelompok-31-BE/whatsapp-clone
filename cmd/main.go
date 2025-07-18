package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "all good",
		})
	})

	log.Println("Server is running on port 8080")

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}

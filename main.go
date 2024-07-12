package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	var address = "127.0.0.1:8080"

	router.GET("/hello/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	log.Fatalln(router.Run(address))
}

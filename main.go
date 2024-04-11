package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting server..")
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	log.Fatal(r.Run())
}

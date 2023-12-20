package main

import (
	"authentication-service/http"
	"authentication-service/repository"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Main function
func main() {
	channel := make(chan string)
	router := gin.Default()
	router.GET("/albums", repository.GetAlbums)
	router.POST("/users", repository.GetUsers)
	router.POST("/event-stream", func(c *gin.Context) {
		fmt.Println("line goes for POST: %s", c.Request.Body)
		http.HandleEventStreamPost(c, channel)
	})
	router.GET("/event-stream", func(c *gin.Context) {
		fmt.Println("line goes for GET: %s", c.Request.Body)
		http.HandleEventStreamGet(c, channel)
	})
	log.Fatalf("error running HTTP server: %s\n", router.Run(":9990"))
}

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Clients is a list of channels to send events to connected clients
var clients = make(map[chan string]struct{})

// broadcast sends an event to all connected clients
func broadcast(data string) {
	for client := range clients {
		client <- data
	}
}

func server() {
	router := gin.Default()

	router.GET("/notifications", func(c *gin.Context) {
		// set the responce header to indicate SEE content type
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// allow all origins to access endpoint
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		// Create a channel to send events to the client
		fmt.Println("Client Connected")
		eventsCh := make(chan string)
		clients[eventsCh] = struct{}{} // add client to clients map
		defer func() {
			delete(clients, eventsCh) // remove the client when they disconnect
			close(eventsCh)
		}()

		// Listen for client close and remove the client from the list
		notify := c.Writer.CloseNotify()
		go func() {
			<-notify
			fmt.Println("Client Disconnected")
		}()

		for {
			data := <-eventsCh
			fmt.Println("Sending data to client", data)
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
		}
	})

	// Handle Post request
	router.POST("/send-data", func(c *gin.Context) {
		data := c.PostForm("data")
		fmt.Println("Data received from the client: ", data)
		broadcast(data)
		c.JSON(http.StatusOK, gin.H{"message": "Data sent to clients"})
	})

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(1 * time.Second)
			broadcast(fmt.Sprint(i))
		}
	}()

	// Start the server
	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}

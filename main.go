package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sendNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("wrong http method type, expected:POST, got:" + r.Method))
		return
	}

	log.Printf("%+v", r)
	w.Write([]byte("Got Notification"))
}

func getNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("wrong http method type expected:POST, got:" + r.Method))
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
		time.Sleep(2 * time.Second)
		w.(http.Flusher).Flush()

	}
}

func main() {
	http.HandleFunc("/send-notification", sendNotification)
	http.HandleFunc("/get-notifications", getNotifications)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Error Starting the Server", err)
	}
}

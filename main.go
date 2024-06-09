package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var localCache map[string][]string

type Notification struct {
	UserName string `json:"sendTo"`
	Data     string `json:"data"`
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func sendNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("wrong http method type, expected:POST, got:" + r.Method))
		return
	}

	var newNotification Notification
	if err := json.NewDecoder(r.Body).Decode(&newNotification); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	userName := newNotification.UserName
	data := newNotification.Data

	log.Println("Received new notification to send")
	log.Printf("UserName: %v, Data: %v", userName, data)

	localCache[userName] = append(localCache[userName], data)

	w.Write([]byte("Got Notification"))
}

func getNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("wrong http method type expected:POST, got:" + r.Method))
		return
	}

	userName := r.URL.Path[len("/get-notifications/"):]
	log.Printf("Requested notification for user: %v", userName)

	notifications, ok := localCache[userName]
	if !ok {
		w.Write([]byte("Don't have notifications"))
		return
	}

	log.Printf("sending notification to user: %+v", notifications[0])
	w.Write([]byte(notifications[0]))

	// w.Header().Set("Content-Type", "text/event-stream")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")

	// for _, v := range localCache["0101"] {
	// 	log.Println(v)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
	// 	time.Sleep(2 * time.Second)
	// 	w.(http.Flusher).Flush()

	// }

	// w.Write([]byte("wait a while"))
}

func main() {
	log.Println("Starting notifcation-service")
	localCache = make(map[string][]string)

	http.HandleFunc("/send-notification", sendNotification)
	http.HandleFunc("/get-notifications/", getNotifications)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Println("Error Starting the Server", err)
	}
}

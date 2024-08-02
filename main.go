package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

type Message struct {
	SendFrom string `json:"sendFrom"`
	SendTo   string `json:"sendTo"`
	Text     string `json:"text"`
}

var readyToSendMsgs map[string][]Message

func readMessages(ch chan<- Message) {
	file, err := os.Open("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optional: resize scanner capacity for lines over 64K
	// const mxCap int = 70000
	// buf := make([]byte, mxCap)
	// scanner.Buffer(buf, mxCap)

	for scanner.Scan() {
		s := scanner.Text()
		msg := Message{}
		json.Unmarshal([]byte(s), &msg)
		ch <- msg
	}

	close(ch)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processMessage(ch <-chan Message) {
	for msg := range ch {
		// log.Printf("Received Message For Processing: %s", msg)
		log.Printf("Received Message For Processing..")
		// Process the message
		rcvMsgs, ok := readyToSendMsgs[msg.SendTo]
		if !ok {
			rcvMsgs = []Message{}
		}
		rcvMsgs = append(rcvMsgs, msg)
		readyToSendMsgs[msg.SendTo] = rcvMsgs
	}
}

func main() {
	log.Println("HELLO :)")

	ch := make(chan Message)
	readyToSendMsgs = make(map[string][]Message)

	go readMessages(ch)

	go processMessage(ch)

	time.Sleep(5 * time.Second)

	for k, _ := range readyToSendMsgs {
		log.Println(k)
	}
}

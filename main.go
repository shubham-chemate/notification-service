package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

type Message struct {
	SendFrom string `json:"sendFrom"`
	SendTo   string `json:"sendTo"`
	Text     string `json:"text"`
}

var readyToSendMsgs map[string][]Message
var servicesList map[string]bool
var usersList map[string]bool
var userPreferences map[string][]string

func readMessages(inputCh chan<- Message) {
	file, err := os.Open("./data-sm.json")
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
		inputCh <- msg
	}

	close(inputCh)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkFormat(msg Message) (bool, string) {
	if msg.SendFrom == "" {
		return false, "Missing Sender"
	}

	if msg.SendTo == "" {
		return false, "Missing Receiver"
	}

	if msg.Text == "" {
		return false, "Missing Text"
	}

	return true, ""
}

func checkSender(msg Message) (bool, string) {
	_, ok := servicesList[msg.SendFrom]
	if !ok {
		return false, "Unknown sender"
	}

	return true, ""
}

func checkReceiver(msg Message) (bool, string) {
	_, ok := usersList[msg.SendTo]
	if !ok {
		return false, "Unknown receiver"
	}

	return true, ""
}

func checkCompliance(msg Message) (bool, string) {
	if strings.Contains(msg.Text, "Gender") || strings.Contains(msg.Text, "gender") {
		return false, "Contains Gender"
	}

	return true, ""
}

func checkUserPref(msg Message) (bool, string) {
	msgType := "xyz"
	for _, v := range userPreferences[msg.SendTo] {
		if msgType == v {
			return true, ""
		}
	}
	return false, "User Preference Not Found"
}

func validateMessage(inputCh <-chan Message, validMsgCh chan<- Message) {
	for msg := range inputCh {
		log.Printf("Received message for validation: %+v", msg)

		ok, failReason := checkFormat(msg)
		if !ok {
			log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
			continue
		}

		ok, failReason = checkSender(msg)
		if !ok {
			log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
			continue
		}

		ok, failReason = checkReceiver(msg)
		if !ok {
			log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
			continue
		}

		ok, failReason = checkCompliance(msg)
		if !ok {
			log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
			continue
		}

		ok, failReason = checkUserPref(msg)
		if !ok {
			log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
			continue
		}

		validMsgCh <- msg
	}
}

func processMessage(validMsgCh <-chan Message, readyToSendCh chan<- Message) {
	for msg := range validMsgCh {
		log.Printf("Received Message For Processing: %v", msg)

		// notification priority

		readyToSendCh <- msg
	}
}

func sendMessage(readyToSendCh <-chan Message) {
	for msg := range readyToSendCh {
		log.Printf("Received Message for Sending: %v", msg)

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

	inputCh := make(chan Message)
	validMsgCh := make(chan Message)
	readyToSendCh := make(chan Message)

	readyToSendMsgs = make(map[string][]Message)
	servicesList = make(map[string]bool)
	usersList = make(map[string]bool)
	userPreferences = make(map[string][]string)

	servicesList["system"] = true
	usersList["Shubham"] = true
	usersList["Rushikesh"] = true
	usersList["Pranav"] = true
	usersList["Kaiwalya"] = true

	userPreferences["Shubham"] = []string{"xyz"}
	userPreferences["Rushikesh"] = []string{"xyz"}
	userPreferences["Pranav"] = []string{"xyz"}
	userPreferences["Kaiwalya"] = []string{"xyz"}

	go readMessages(inputCh)

	go validateMessage(inputCh, validMsgCh)

	go processMessage(validMsgCh, readyToSendCh)

	go sendMessage(readyToSendCh)

	time.Sleep(5 * time.Second)

	for k, v := range readyToSendMsgs {
		log.Println(k, v)
	}
}

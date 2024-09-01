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
var userMsgCnt map[string]int
var MAX_MSGS_PER_DAY = 1

func readMessages(inputCh chan Message) {
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

func processMessages(inputCh <-chan Message) {
	for msg := range inputCh {
		go validateMessage(msg)
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

func checkUserMsgCnt(msg Message) (bool, string) {
	cnt, ok := userMsgCnt[msg.SendTo]
	if !ok {
		userMsgCnt[msg.SendTo] = 1
		return true, ""
	}

	if cnt < MAX_MSGS_PER_DAY {
		userMsgCnt[msg.SendTo] = cnt + 1
		return true, ""
	}

	return false, "User already exceeded limit for number of per day messages"
}

func validateMessage(msg Message) {
	log.Printf("Received message for validation: %+v", msg)

	ok, failReason := checkFormat(msg)
	if !ok {
		log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
		return
	}

	ok, failReason = checkSender(msg)
	if !ok {
		log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
		return
	}

	ok, failReason = checkReceiver(msg)
	if !ok {
		log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
		return
	}

	ok, failReason = checkCompliance(msg)
	if !ok {
		log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
		return
	}

	ok, failReason = checkUserPref(msg)
	if !ok {
		log.Printf("Validation Failed, msg: %+v, error: %v", msg, failReason)
		return
	}

	ok, failReason = checkUserMsgCnt(msg)
	if !ok {
		log.Printf("Validation Failed, msg: %v, error: %v", msg, failReason)
		return
	}

	processMessage(msg)
}

func processMessage(msg Message) {
	log.Printf("Received Message For Processing: %v", msg)

	// store the msg in the db for analysis purpose
	// sync store for low-priority msgs
	// async store for high-priority msgs
	// write heavy db

	sendMessage(msg)
}

func sendMessage(msg Message) {
	log.Printf("Received Message for Sending: %v", msg)

	rcvMsgs, ok := readyToSendMsgs[msg.SendTo]
	if !ok {
		rcvMsgs = []Message{}
	}
	rcvMsgs = append(rcvMsgs, msg)
	readyToSendMsgs[msg.SendTo] = rcvMsgs
}

func main1() {
	log.Println("HELLO :)")

	inputCh := make(chan Message)

	readyToSendMsgs = make(map[string][]Message)
	servicesList = make(map[string]bool)
	usersList = make(map[string]bool)
	userPreferences = make(map[string][]string)
	userMsgCnt = make(map[string]int)

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
	go processMessages(inputCh)

	time.Sleep(5 * time.Second)

	for k, v := range readyToSendMsgs {
		log.Println(k, v)
	}
}

func main() {
	// ih := &IntHeap{}
	// heap.Init(ih)
	// heap.Push(ih, 1)
	// heap.Push(ih, 5)
	// heap.Push(ih, 2)
	// heap.Push(ih, 4)

	// fmt.Println(heap.Pop(ih))
	// fmt.Println(heap.Pop(ih))

	// heap.Push(ih, 1)
	// fmt.Println(heap.Pop(ih))

	sh := SafeHeap{}

	sh.Push(1)
	sh.Push(5)
	sh.Push(3)
	sh.Push(2)

	log.Println(sh.Pop())
	log.Println(sh.Pop())

	// sh.Push(3)
	// log.Println(sh.Pop())
	// log.Println(sh.Pop())

}

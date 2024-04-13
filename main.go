package main

import (
	"fmt"
	"math/rand"
	n "notification-service/notification"
	"time"
)

func floodNotifications(vhpnCh chan<- n.Notification, nCh chan<- n.Notification) {
	for i := 0; i < 100; i++ {
		randMS := rand.Intn(100)
		nn := n.Notification{
			Content:  fmt.Sprintf("Notification Number: %d", i),
			Priority: "very-high",
			SendAt:   time.Now().Add(time.Duration(randMS) * time.Second),
		}
		if i%5 == 0 {
			vhpnCh <- nn
		} else {
			if i%2 == 2 {
				nn.Priority = "medium"
			} else if i%3 == 0 {
				nn.Priority = "high"
			} else {
				nn.Priority = "low"
			}
			nCh <- nn
		}
	}
}

func handleNotification(nn n.Notification, rnCh chan n.Notification) {
	time.Sleep(time.Until(nn.SendAt))
	// sending the notification
	fmt.Println("Notfication Send Successfully")
	fmt.Println(nn.Content, nn.SendAt.Format("03:04:05PM"), nn.Priority)
	nn.SendAttempts += 1

	failed := false
	if failed {

		// 1 initial send attmpt & 3 retries = total 4
		if nn.Priority == "very-high" && nn.SendAttempts < 4 {
			rnCh <- nn
		}

		if nn.Priority == "high" && nn.SendAttempts < 3 {
			rnCh <- nn
		}

		if nn.Priority == "medium" && nn.SendAttempts < 2 {
			rnCh <- nn
		}
	}
}

func main() {

	vhpnCh := make(chan n.Notification)
	nCh := make(chan n.Notification)
	rnCh := make(chan n.Notification)

	go floodNotifications(vhpnCh, nCh)

	for i := 0; i < 100; i++ {
		if nn, ok := <-vhpnCh; ok {
			go handleNotification(nn, rnCh)
			continue
		}
		if nn, ok := <-nCh; ok {
			go handleNotification(nn, rnCh)
			continue
		}
		if rn, ok := <-rnCh; ok {
			go handleNotification(rn, rnCh)
		}
	}

}

package main

import (
	"fmt"
	"math/rand"
	n "notification-service/notification"
	"time"
)

// type VeryHighPriorityNQ struct {
// 	nq nq.NQ
// }

// type NotificationQ struct {
// 	nq nq.NQ
// }

// type HighPriorityNQ struct {
// 	nq nq.NQ
// }

// type MediumPriorityNQ struct {
// 	nq nq.NQ
// }

// type LowPriorityNQ struct {
// 	nq nq.NQ
// }

// type RetryNQ []*RetryNotification

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

func handleNotification(nn n.Notification) {
	time.Sleep(time.Until(nn.SendAt))
	// sending the notification
	fmt.Println("Notfication Send Successfully")
	fmt.Println(nn.Content, nn.SendAt.Format("03:04:05PM"), nn.Priority)
}

func main() {

	vhpnCh := make(chan n.Notification)
	nCh := make(chan n.Notification)

	go floodNotifications(vhpnCh, nCh)

	// vhpnQ := VeryHighPriorityNQ{
	// 	nq: make(nq.NQ, 0),
	// }
	// heap.Init(&vhpnQ.nq)

	// nQ := NotificationQ{
	// 	nq: make(nq.NQ, 0),
	// }
	// heap.Init(&nQ.nq)

	for i := 0; i < 100; i++ {
		select {
		case nn := <-vhpnCh:
			go handleNotification(nn)
		case nn := <-nCh:
			go handleNotification(nn)
		}
	}

	select {}

	// s, err := gocron.NewScheduler()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fn := func(s string) {
	// 	log.Println(s)
	// }

	// j, err := s.NewJob(gocron.DurationJob(10*time.Second), gocron.NewTask(fn, "first-job"))
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// log.Println(j.ID())

	// j, err = s.NewJob(gocron.DurationJob(13*time.Second), gocron.NewTask(fn, "second-job"))
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// log.Println(j.ID())

	// tm, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Apr 12, 2024 at 11:50am (IST)")
	// s.NewJob(
	// 	gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(tm)),
	// 	gocron.NewTask(fn, "third-job"),
	// )

	// s.Start()

	// select {}

	// s.Shutdown()

	// vhpnq := VeryHighPriorityNQ{}
	// hpnq := HighPriorityNQ{}
	// mpnq := MediumPriorityNQ{}
	// lpnq := LowPriorityNQ{}

	// vhpnq.nq = make(nq.NQ, 0)
	// heap.Init(&vhpnq.nq)

	// hpnq.nq = make(nq.NQ, 0)
	// heap.Init(&hpnq.nq)

	// mpnq.nq = make(nq.NQ, 0)
	// heap.Init(&mpnq.nq)

	// lpnq.nq = make(nq.NQ, 0)
	// heap.Init(&lpnq.nq)
	// heap.Push(&hpnq.nq, notif2)
	// heap.Push(&hpnq.nq, notif4)
	// heap.Push(&hpnq.nq, notif3)
	// heap.Push(&hpnq.nq, notif1)
	// heap.Push(&hpnq.nq, notif5)

	// for hpnq.nq.Len() > 0 {
	// 	notif := heap.Pop(&hpnq.nq).(*Notification)
	// 	fmt.Println(notif)
	// }

	// veryHighPriorityNCh := make(chan n.Notification, 5000)
	// highPriorityNCh := make(chan n.Notification, 5000)
	// medPriorityNCh := make(chan n.Notification, 5000)
	// lowPriorityNCh := make(chan n.Notification, 5000)

	// for {
	// 	notif, ok := <-veryHighPriorityNCh
	// 	if ok {
	// 		heap.Push(&vhpnq.nq, notif)
	// 		continue
	// 	}

	// 	notif, ok = <-highPriorityNCh
	// 	if ok {
	// 		heap.Push(&hpnq.nq, notif)
	// 		continue
	// 	}

	// 	notif, ok = <-medPriorityNCh
	// 	if ok {
	// 		heap.Push(&mpnq.nq, notif)
	// 		continue
	// 	}

	// 	notif, ok = <-lowPriorityNCh
	// 	if ok {
	// 		heap.Push(&lpnq.nq, notif)
	// 		continue
	// 	}

	// 	// check for retry queue
	// }

	// timer := time.NewTimer(1 * time.Second)
	// defer timer.Stop()

	// tm, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Apr 10, 2024 at 4:16pm (IST)")

	// var wg sync.WaitGroup
	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		ii := time.Duration(100 - i)
	// 		time.Sleep(ii * time.Millisecond)
	// 		fmt.Printf("GR %d: %v\n", i, time.Now())
	// 	}(i)
	// }
	// wg.Wait()

	// for i := 0; i < 60; i++ {
	// <-timer.C
	// fmt.Println(time.Now())
	// }

}

package main

import (
	"log"
	"notification-service/nq"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type VeryHighPriorityNQ struct {
	nq nq.NQ
}

type HighPriorityNQ struct {
	nq nq.NQ
}

type MediumPriorityNQ struct {
	nq nq.NQ
}

type LowPriorityNQ struct {
	nq nq.NQ
}

// type RetryNQ []*RetryNotification

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Println(err)
		return
	}

	j, err := s.NewJob(gocron.DurationJob(
		10*time.Second,
	),
		gocron.NewTask(func(s string) {
			log.Println(s)
		}, "first-job"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(j.ID())

	s.Start()

	select {}

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

package main

import (
	"fmt"
	"notification-service/nq"
	"time"
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

	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	for i := 0; i < 60; i++ {
		<-timer.C
		fmt.Println(time.Now())
	}

}

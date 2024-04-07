package main

import (
	"container/heap"
	"fmt"
)

type Notification struct {
	Content  string
	SendAt   int
	Priority string
}

type HighPriorityNotificationQueue []*Notification

func (pq HighPriorityNotificationQueue) Len() int { return len(pq) }

func (pq HighPriorityNotificationQueue) Less(i, j int) bool {
	return pq[i].SendAt < pq[j].SendAt
}

func (pq HighPriorityNotificationQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *HighPriorityNotificationQueue) Push(notif any) {
	newNotif := notif.(Notification)
	*pq = append(*pq, &newNotif)
}

func (pq *HighPriorityNotificationQueue) Pop() any {
	old := *pq
	n := len(old)
	poppedNotif := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return poppedNotif
}

func main() {
	// notif := Notification{
	// 	Content: "OTP",
	// 	Priority: "very-high",
	// 	SendAt:
	// }

	// h := &IntHeap{2, 1, 5}
	// heap.Init(h)
	// heap.Push(h, 6)
	// heap.Push(h, 0)
	// fmt.Printf("minimum: %d\n", (*h)[0])
	// for h.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(h))
	// }
	// fmt.Println()

	// items := map[string]int{
	// 	"banana": 3,
	// 	"apple":  2,
	// 	"pear":   4,
	// }

	// pq := make(PQ, len(items))
	// i := 0
	// for v, p := range items {
	// 	pq[i] = &Item{
	// 		Val: v,
	// 		P:   p,
	// 		// Ind: i,
	// 	}
	// 	i++
	// }

	// heap.Init(&pq)

	// item := &Item{
	// 	Val: "orange",
	// 	P:   1,
	// }
	// heap.Push(&pq, item)
	// // pq.update(item, item.value, 5)
	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*Item)
	// 	fmt.Printf("%.2d:%s ", item.P, item.Val)
	// }

	// pq := PriorityQueue{}
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second)
	// 	notif := Notification{
	// 		CustomerId: i + 1,
	// 		SendAt:     time.Now(),
	// 	}
	// 	pq = append(pq, &notif)
	// }

	// tm1, _ := time.Parse("03:04:05PM", "03:35:00PM")
	notif1 := Notification{Content: "OTP: 1234", SendAt: 1}

	// tm2, _ := time.Parse("03:04:05PM", "03:34:00PM")
	notif2 := Notification{Content: "OTP: 3455", SendAt: 2}

	// tm3, _ := time.Parse("03:04:05PM", "03:33:00PM")
	notif3 := Notification{Content: "OTP: 1234", SendAt: 3}

	// tm4, _ := time.Parse("03:04:05PM", "03:36:00PM")
	notif4 := Notification{Content: "OTP: 4523", SendAt: 4}

	// tm5, _ := time.Parse("03:04:05PM", "03:37:00PM")
	notif5 := Notification{Content: "Money Sent: INR 200", SendAt: 5}

	pq := make(HighPriorityNotificationQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, notif2)
	heap.Push(&pq, notif4)
	heap.Push(&pq, notif3)
	heap.Push(&pq, notif1)
	heap.Push(&pq, notif5)

	for pq.Len() > 0 {
		notif := heap.Pop(&pq).(*Notification)
		fmt.Println(notif)
	}

	// pq.Push(&notif1)
	// pq.Push(&notif2)
	// pq.Push(&notif3)
	// pq.Push(&notif4)
	// pq.Push(&notif5)

	// for i := 0; i < 5; i++ {
	// pq.Pop()
	// for _, q := range pq {
	// 	log.Println(q.SendAt)
	// }
	// }
}

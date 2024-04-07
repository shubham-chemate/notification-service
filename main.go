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
	notif1 := Notification{Content: "OTP: 1234", SendAt: 1}
	notif2 := Notification{Content: "OTP: 3455", SendAt: 2}
	notif3 := Notification{Content: "OTP: 1234", SendAt: 3}
	notif4 := Notification{Content: "OTP: 4523", SendAt: 4}
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
}

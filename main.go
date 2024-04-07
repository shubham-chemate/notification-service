package main

import (
	"container/heap"
	"fmt"
)

// type Notification struct {
// CustomerId int
// Content    string
// SendAt time.Time
// Priority   string
// }

// https://pkg.go.dev/container/heap
// type PriorityQueue []*Notification

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].SendAt.Before(pq[j].SendAt)
// }

// func (pq PriorityQueue) Swap(i, j int) {

// }

// func (pq *PriorityQueue) Push(x any) {
// 	notification := x.(*Notification)
// 	*pq = append(*pq, notification)
// }

// func (pq *PriorityQueue) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	notif := old[n-1]
// 	old[n-1] = nil
// 	*pq = old[:n-1]
// 	return notif
// }

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 6)
	heap.Push(h, 0)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()

	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}

	pq := make(PQ, len(items))
	i := 0
	for v, p := range items {
		pq[i] = &Item{
			Val: v,
			P:   p,
			// Ind: i,
		}
		i++
	}

	heap.Init(&pq)

	item := &Item{
		Val: "orange",
		P:   1,
	}
	heap.Push(&pq, item)
	// pq.update(item, item.value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.P, item.Val)
	}

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
	// notif1 := Notification{CustomerId: 1, SendAt: tm1}

	// tm2, _ := time.Parse("03:04:05PM", "03:34:00PM")
	// notif2 := Notification{CustomerId: 2, SendAt: tm2}

	// tm3, _ := time.Parse("03:04:05PM", "03:33:00PM")
	// notif3 := Notification{CustomerId: 3, SendAt: tm3}

	// tm4, _ := time.Parse("03:04:05PM", "03:36:00PM")
	// notif4 := Notification{CustomerId: 4, SendAt: tm4}

	// tm5, _ := time.Parse("03:04:05PM", "03:37:00PM")
	// notif5 := Notification{CustomerId: 5, SendAt: tm5}

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

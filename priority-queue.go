package main

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Item struct {
	Val string
	P   int
	// Ind int
}

type PQ []*Item

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].P > pq[j].P
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	// pq[i].Ind = i
	// pq[j].Ind = j
}

func (pq *PQ) Push(x any) {
	// n := len(*pq)
	item := x.(*Item)
	// item.Ind = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	// item.Ind = -1
	*pq = old[0 : n-1]
	return item
}

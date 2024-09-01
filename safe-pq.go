package main

import "log"

type SafeHeap []int

func (h SafeHeap) swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}

func (h SafeHeap) less(pi, ci int) bool {
	return h[pi] > h[ci]
}

func (h *SafeHeap) up(i int) {
	for {
		p := (i - 1) / 2
		if p == i || !h.less(p, i) {
			break
		}
		h.swap(p, i)
		i = p
	}
}

func (h SafeHeap) isSmall(x, y int) bool {
	return h[x] < h[y]
}

func (h *SafeHeap) down(i int) {
	n := len(*h)
	for {
		lc := 2*i + 1
		if lc >= n || lc < 0 {
			break
		}
		ci := lc
		if rc := 2*i + 2; rc < n && h.isSmall(rc, lc) {
			ci = rc
		}
		if h.isSmall(i, ci) {
			break
		}
		h.swap(i, ci)
		i = ci

	}
}

func (h *SafeHeap) Push(x int) {
	*h = append(*h, x)
	log.Println(*h)
	h.up(len(*h) - 1)
	log.Println(*h)
}

func (h *SafeHeap) Pop() int {
	old := *h
	n := len(old)

	x := old[0]
	old.swap(0, n-1)
	old = old[0 : n-1]
	old.down(0)
	*h = old

	return x
}

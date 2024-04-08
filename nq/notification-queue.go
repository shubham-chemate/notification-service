package nq

import n "notification-service/notification"

type NQ []*n.Notification

func (pq NQ) Len() int { return len(pq) }

func (pq NQ) Less(i, j int) bool {
	return pq[i].SendAt < pq[j].SendAt
}

func (pq NQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *NQ) Push(notif any) {
	newNotif := notif.(n.Notification)
	*pq = append(*pq, &newNotif)
}

func (pq *NQ) Pop() any {
	old := *pq
	n := len(old)
	poppedNotif := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return poppedNotif
}

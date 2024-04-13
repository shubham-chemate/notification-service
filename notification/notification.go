package notification

import "time"

type Notification struct {
	Content      string
	SendAt       time.Time
	Priority     string
	SendAttempts int
}

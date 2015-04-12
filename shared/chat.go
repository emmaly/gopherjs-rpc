package shared

import "time"

// ChatMessage is a chat message
type ChatMessage struct {
	Name    string
	Time    time.Time
	Message string
}

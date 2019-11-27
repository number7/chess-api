package structs

import "time"

type Heartbeat struct {
	Message   string
	Timestamp time.Time
}

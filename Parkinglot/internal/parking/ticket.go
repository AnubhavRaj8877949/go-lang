package parking

import "time"

type Ticket struct {
	TicketID  string
	SpotID    string
	Floor     int
	Vehicle   Vehicle // copy (pointer nahi) — bahar se koi badal na sake
	EntryTime time.Time
}

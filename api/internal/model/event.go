package model

import "time"

type Event struct {
	Name    string
	StartAt time.Time
	EndAt   time.Time
}

func (e Event) IsInSession(now time.Time) bool {
	// including StartAt but not EndAt (see examples in test code)
	return e.EndAt.After(now) && !e.StartAt.After(now)
}

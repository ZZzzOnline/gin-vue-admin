package norm

import "time"

type DeletedAt NullTime

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

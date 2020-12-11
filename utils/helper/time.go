package helper

import "time"

/*
	OnlyGetCurrentTime:
		* to get current time in one place
 */
func OnlyGetCurrentTime() (time.Time) {
	return time.Now()
}

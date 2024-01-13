package utils

import (
	"fmt"
	"time"
)

func SameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func LogPrettyTime() {
	now := time.Now()
    formattedTime := now.Format("Monday, 02-Jan-2006 15:04:05\n")
    fmt.Print(formattedTime)
}

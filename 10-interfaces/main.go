package main

import (
	"fmt"
	"time"
)

type clock interface {
	now() time.Time
}

type realClock struct{}

func (realClock) now() time.Time {
	return time.Now()
}

func isFriday(c clock) bool {
	weekday := c.now().Weekday()
	return weekday == time.Friday
}

func main() {
	if isFriday(&realClock{}) {
		fmt.Println("Have a nice weekend!")
	} else {
		fmt.Println("It's not Friday yet :-(")
	}
}

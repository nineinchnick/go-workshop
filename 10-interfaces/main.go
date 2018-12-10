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
	if !isRealClock(c) {
		fmt.Println("Warning! Not using a realClock")
	}
	return weekday == time.Friday
}

func isRealClock(c interface{}) bool {
	_, ok := c.(*realClock)
	return ok
}

func main() {
	if isFriday(&realClock{}) {
		fmt.Println("Have a nice weekend!")
	} else {
		fmt.Println("It's not Friday yet :-(")
	}
}

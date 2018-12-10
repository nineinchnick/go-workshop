package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	linux = "Linux"
)

func printOS() {
	fmt.Print("Go runs on ")
	// switch can compare anything
	switch os := runtime.GOOS; strings.ToLower(os) {
	case "darwin":
		fmt.Println("OS X.")
		// no break needed
	case strings.ToLower(linux):
		fmt.Println(linux)
	default:
		fmt.Printf("%s.", os)
	}
}

func getMonth() string {
	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	defer fmt.Println("Done detecting current month")

	result := "unknown"

	defer func() {
		fmt.Println("result is: ", result)
	}()

	switch time.Now().Month() {
	case 1:
		result = "January"
	case 2:
		result = "February"
	default:
		result = "No way I'm doing all of them"
	}
	return result
}

package main

import (
	"errors"
	"fmt"
)

// convert to named return variable
func sum(a int, b int) (*int, error) {
	if a < 0 || b < 0 {
		return nil, errors.New("cannot add negative numbers")
	}
	result := a + b
	return &result, nil
}

func main() {
	result, err := sum(5, -2)
	fmt.Println("5 - 2 =", result, err)

	c := func(a string) {
		fmt.Println("'a' is a copy but result is available in a closure: %d", result)
	}
	c("something")
}

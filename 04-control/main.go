package main

import (
	"fmt"
)

func main() {
	// simplest loop
	a := 0
	for i := 0; i < 10; i++ {
		a += i
	}
	fmt.Println("a =", a)

	// look ma', no while!
	b := 1
	for b < 100 {
		b += b
	}
	fmt.Println("b =", b)

	// infinite?
	c := 1
	for {
		c += c
		if c < 100 {
			continue
		}
		// skip 'else' if the condition above stops execution
		break
	}

	printOS()
	fmt.Println("Current month: ", getMonth())
}

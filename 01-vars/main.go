package main

import (
	"fmt"
)

const (
	// types can be inferred during assignment
	a = "a"
)

// all types have zero-value
var b int

var (
	c
	d = 0.1
	e = `{"wow": "such",
"json": "!"}`
)

func main() {
	local := 5
	fmt.Printf("All my consts and vars: %s, %d, %s, %f, %s\n", a, b, c, d, e)
	fmt.Printf("Details: %v, %#v\n", local, local)
}

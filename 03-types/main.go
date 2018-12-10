package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe    bool       = false
	maxInt             = 1<<64 - 1
	maxUInt uint64     = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", maxUInt, maxUInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversion
	fmt.Printf("Is %v greater than %v? %b", maxInt, maxUInt, maxInt > maxUInt)
}

package main

import "fmt"

type point struct {
	x, y int
}

var (
	p1  = point{1, 2}
	p2  = point{x: 1}  // y:0 is implicit
	p3  = point{}      // x:0 and y:0
	ptr = &point{1, 2} // has type *point

)

func main() {
	fmt.Println(p1, p2, p3, ptr)

	p1.x = 4
	(*ptr).x = 5
	ptr.x = 6

	fmt.Println(p1, p2, p3, ptr)

	p4 := struct{ i, j, k int }{1, 2, 3}
	fmt.Println(p4)
}

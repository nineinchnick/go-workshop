package main

import "fmt"

type point struct {
	x int
	y int
}

// not inheritance, but embedding: https://golang.org/doc/effective_go.html#embedding
type point2 struct {
	point
	y string
}

var (
	p1  = point{1, 2}
	p2  = point{y: 5, x: 1} // y:0 is implicit
	p3  = point{}           // x:0 and y:0
	ptr = &point{1, 2}      // has type *point

)

func printPoint(p *point) {
	fmt.Println(p)
}

func main() {
	fmt.Println(p1, p2, p3, ptr)

	x := point2{}
	x.point.x = 5
	x.x = 5
	fmt.Println("x: ", x)

	p1.x = 4
	(*ptr).x = 5
	ptr.x = 6

	fmt.Println(p1, p2, p3, ptr)

	p4 := []struct{ i, j, k int }{
		{1, 2, 3},
	}
	for _, pp := range p4 {
		fmt.Println(pp)
	}

	printPoint(&point{1, 2})
}

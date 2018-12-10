package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p point) abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p point) move(angle, distance float64) {
	p.x += math.Cos(angle) * distance
	p.y += math.Sin(angle) * distance
}

func main() {
	a := point{3, 5}
	fmt.Println("Distance of", a, "from center:", a.abs())
	a.move(45.0, 15.0)
	fmt.Println("Moving by 15 at 45 degrees")
	fmt.Println("Distance of", a, "from center:", a.abs())

	// methods are just functions but there's automatic pointer dereference for convienience
}

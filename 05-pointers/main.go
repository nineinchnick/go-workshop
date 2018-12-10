package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Printf("Starting with i=%d, j=%d\n", i, j)

	var p *int
	fmt.Println("Empty pointer:", p)

	p = &i                                  // point to i
	fmt.Println("Dereference p:", *p, p)    // read i through the pointer
	*p = 21                                 // set i through the pointer
	fmt.Println("Assigned i through p:", i) // see the new value of i

	p = &j                                        // point to j
	*p = *p / 37                                  // divide j through the pointer
	fmt.Println("Divided j by 37, through p:", j) // see the new value of j

	// no pointer arithmetic!
	c := p + &i
}

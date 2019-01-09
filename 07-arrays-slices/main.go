package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	primes := []int{2, 3, 5, 7, 11, 13}

	fmt.Printf("a: %#v\na[1:]: %#v\nprimes: %#v\n", a, a[1:], primes)

	var morePrimes []int
	fmt.Printf("morePrimes: %#v\n", morePrimes)

	morePrimes = append(morePrimes, 17)
	fmt.Printf("append to nil slice: %#v\n", morePrimes)

	morePrimes = make([]int, 0)
	morePrimes = append(morePrimes, 17, 19)
	fmt.Printf("append after make: %#v\n", morePrimes)

	var primesCopy []int
	for _, p := range morePrimes {
		primesCopy = append(primesCopy, p)
	}

	// the _only_ single line trick in Go...
	copy := append([]int{}, morePrimes...)
	fmt.Printf("copy: %#v\n", copy)

	// iterate
	fmt.Println("loop 1")
	for i, n := range copy {
		fmt.Printf("item %d at index %d\n", n, i)
	}

	fmt.Println("loop 2")
	for _, n := range copy {
		fmt.Printf("item %d\n", n)
	}

	fmt.Println("loop 3")
	for i := range copy {
		fmt.Printf("index %d\n", i)
	}
}

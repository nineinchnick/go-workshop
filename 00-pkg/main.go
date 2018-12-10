package main

import mypkg "github.com/nineinchnick/go-workshop/sum"

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.intn(10))
	fmt.Println("2 + 2 =", mypkg.Sum(2, 2))
}

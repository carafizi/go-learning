package main

import (
	"fmt" 
	"math/rand"
)
var k = 3

func add(x, y int) int {
	return x + y
}

func main() {
	z := 4
	fmt.Println(k, z)
	fmt.Println("My favorite number is", rand.Intn(100))	
	fmt.Println(add(42, 13))
}
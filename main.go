package main

import (
	"fmt"
	"math/bits"
)

func sum(a, b int) int {
	sum := a + b
	return sum
}

func subtract (a , b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}

}

func multiply (a , b int) int {
	return a * b

}

func main() {
	total := sum(2, 2)
	subtract := subtract(3,1)
	multi := multiply(3,1)
	fmt.Println("Hello world!")
	fmt.Println(total)
	fmt.Println(subtract)
	fmt.Println(multi)

}

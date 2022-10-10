package main

import (
	"fmt"
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


func main() {
	total := sum(2, 2)
	subtract := subtract(3,1)
	fmt.Println("Hello world!")
	fmt.Println(total)
	fmt.Println(subtract)

}

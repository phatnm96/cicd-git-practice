package main

import (
	"fmt"
)

func sum(a, b int) int {
	sum := a + b
	return sum
}


func main() {
	total := sum(2, 2)
	fmt.Println("Hello world!")
	fmt.Println(total)

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	defer fmt.Println("D")
	// A D C B

	for i := 0; i < 3; i++ {
		// 2 1 0
		defer fmt.Println(i)
		defer func() {
			// 3 3 3
			fmt.Println(i)
		}()
	}
}

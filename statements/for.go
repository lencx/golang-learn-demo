package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		// 0 - 9
		fmt.Println(&i, i)
	}

	a := 1
	for {
		a++
		if a > 5 {
			return
		}
		// 1 - 4
		fmt.Println(a)
	}
	// OR
	// for a < 5 {
	// 	fmt.Println(a)
	// 	a++
	// }
}

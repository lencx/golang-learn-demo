package main

import "fmt"

// 6: 0110
// 11: 1011
func main() {
	// 2
	fmt.Println(6 & 11)
	// 15
	fmt.Println(6 | 11)
	// 13
	fmt.Println(6 ^ 11)
	// 4
	fmt.Println(6 &^ 11)
}

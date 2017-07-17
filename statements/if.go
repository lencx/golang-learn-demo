package main

import "fmt"

func main() {

	a := 5
	// a := 8; is initialize the statement
	if a := 8; a > 6 {
		// 8
		fmt.Println(a)
	}

	// a is 5
	if a == 5 {
		fmt.Println("a is", a)
	}

	// At the same time declare multiple variables
	if c, d := 3, 7; c > d {
		fmt.Println(c, ">", d)
	} else {
		// 3 < 7
		fmt.Println(c, "<", d)
	}
}

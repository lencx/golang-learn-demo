package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a is 0")
	case 1:
		fmt.Println("a is 1")
	default:
		fmt.Println("a is", a)
	}

	switch {
	case a >= 0:
		fmt.Println("a >= 0")
	case a < 0:
		fmt.Println("a < 0")
	}

	switch b := 5; {
	case b >= 0:
		fmt.Println("b >= 0")
		fallthrough
	case b == 5:
		fmt.Println("b == 5")
	case b < 0:
		fmt.Println("b < 0")
	}
}

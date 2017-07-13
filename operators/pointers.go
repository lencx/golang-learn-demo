package main

import "fmt"

func main() {
	a := 2
	var p *int = &a
	// 0xc420072178
	fmt.Println(p)
	// 2
	fmt.Println(*p)
}

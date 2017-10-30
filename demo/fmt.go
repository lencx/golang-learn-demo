package main

import (
	"fmt"
)

func main() {
	var i interface{} = 45
	num := 100
	isShow := true
	pi := 3.1415926
	fmt.Printf("%v\n", i)      // 45
	fmt.Printf("%b\n", num)    // 1100100
	fmt.Printf("%c\n", num)    // d
	fmt.Printf("%d\n", num)    // 100
	fmt.Printf("%e\n", pi)     // 3.141593e+00
	fmt.Printf("%.3f\n", pi)   // 3.142
	fmt.Printf("%o\n", num)    // 144
	fmt.Printf("%T\n", isShow) // bool
	fmt.Printf("%t\n", isShow) // true
	fmt.Printf("%x\n", num)    // 64
}

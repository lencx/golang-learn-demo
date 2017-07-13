package main

import (
	"fmt"
)

const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// 1024
	fmt.Println(KB)
	// 1.048576e+06
	fmt.Println(MB)
	// 1.073741824e+09
	fmt.Println(GB)
	// 1.099511627776e+12
	fmt.Println(TB)
	// 1.125899906842624e+15
	fmt.Println(PB)
	// 1.152921504606847e+18
	fmt.Println(EB)
	// 1.1805916207174113e+21
	fmt.Println(ZB)
	// 1.2089258196146292e+24
	fmt.Println(YB)
}

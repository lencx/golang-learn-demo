package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var c, js, python, rust bool

// m := 4

var (
	ToBe   bool       = false
	Maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	var i int
	python := true
	fmt.Println(i, c, js, python, rust) // 0 false false true false

	var word, word2 string = "SXWICALH", "Len"
	_firstChar := string([]rune(word)[0])
	fmt.Println(_firstChar) // S

	var j = 0
	for j < len(word2) {
		a := string([]rune(word2)[j])
		fmt.Println(a) // L e n
		j += 1
	}

	fmt.Println("-----------------")

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)     // bool(false)
	fmt.Printf(f, Maxint, Maxint) // uint64(18446744073709551615)
	fmt.Printf(f, z, z)           // bool(false)

	fmt.Println("-----------------")

	var _i int     // 0
	var _f float64 // 0
	var _b bool    // false
	var _s string  // ""

	fmt.Printf("%v %v %v %v \n", _i, _f, _b, _s)

	fmt.Println("-----------------")
	var _x, _y int = 3, 4
	var _f2 = math.Sqrt(float64(_x*_x + _y*_y))
	// var _z2 int = int(_f2)
	_z2 := int(_f2)

	_x2 := math.Pow(4, 2)

	fmt.Println(_f2, _z2, _x2) // 5 5 16
	fmt.Printf(f, _f2, _f2)    // float64(5)
	fmt.Printf(f, _z2, _z2)    // int(5)

	v2 := 43
	// v2 is of type int
	fmt.Printf("v2 is of type %T\n", v2)
	v2 = 40
	fmt.Println(v2) // 40
	_v2 := 43.4
	v3 := 0.435 + 0.5i
	// _v2 is of type float64
	fmt.Printf("_v2 is of type %T\n", _v2)
	// v3 is of type complex128
	fmt.Printf("v3 is of type %T\n", v3)

	const Hello = "ni hao"
	fmt.Printf("hello, %s\n", Hello) // hello, ni hao
}

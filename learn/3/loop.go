package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		fmt.Println(math.Abs(tmp - z))

		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := float64(1)
	var tmp float64
	for i := 0; i < 10; i++ {
		tmp = z - (z*z-x)/(2*z)
		z = tmp
		fmt.Println(z)
	}
	return z
}

func main() {
	// fmt.Println(Sqrt(2.0))
	fmt.Println(Sqrt2(2.0))
}

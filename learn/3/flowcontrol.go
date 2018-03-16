package main

import (
	"fmt"
	"math"
)

func main() {
	forFn()
	fmt.Println(sqrt(3), sqrt(-2))
	fmt.Println(pow(3, 3, 5))  // 5
	fmt.Println(pow(3, 2, 10)) // 9
}

func forFn() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum) // 45
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		// fmt.Printf("v: %f", v)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// fmt.Printf("lim: %f", lim)
	return lim
}

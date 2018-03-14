package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(50))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		a := rand.Intn(50)
		fmt.Println("Random: ", a)
	}

	fmt.Println("Now you have %g problems.", math.Nextafter(2, 3))

	fmt.Println(math.Pi)

	fmt.Println(add(4, 8))     // 12
	fmt.Println(add2(4, 8, 3)) // 15

	a, b, c := swap("hello", "len", 24)
	fmt.Println(a, b, c) // hello len 24

	fmt.Println(split(8)) // 80 -72
}

func swap(x, y string, z int32) (string, string, int32) {
	return x, y, z
}

func split(sum int) (x, y int) {
	x = sum * 10
	y = sum - x
	return
}

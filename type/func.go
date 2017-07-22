package main

import (
	"fmt"
)

func main() {
	// 1 3 5
	fn1()
	// [2 4 6 8 10]
	fn2(2, 4, 6, 8, 10)

	/* value copy */
	a, b, c := 1, 4, 7
	// [1 8 9]
	fn3(a, b, c)
	// 1 4 7
	fmt.Println(a, b, c)

	/* address copy */
	s := []int{5, 6, 7, 8, 9}
	// [5 6 6 8 9]
	fn4(s)
	// [5 6 6 8 9]
	fmt.Println(s)

	/* pointers */
	d := 99
	// 2
	fn5(&d)
	// 2
	fmt.Println(d)

	/* anonymous function */
	e := func() {
		fmt.Println("Print E")
	}
	e()

	/* closure */
	f := closure(4)
	// 20
	fmt.Println(f(5))
	// fmt.Println(closure(4)(5))
}

func fn1() (a, b, c int) {
	a, b, c = 1, 3, 5
	fmt.Println(a, b, c)
	return a, b, c
}

func fn2(a ...int) {
	fmt.Println(a)
}

func fn3(s ...int) {
	s[1] = 8
	s[2] = 9
	fmt.Println(s)
}

func fn4(s []int) {
	s[1] = 6
	s[2] = 6
	fmt.Println(s)
}

func fn5(a *int) {
	*a = 2
	fmt.Println(*a)
}

func closure(x int) func(int) int {
	// 0xc420072290
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		// 0xc420072290
		fmt.Printf("%p\n", &x)
		return x * y
	}
}

package main

import "fmt"

func main() {
	arr()
	slice()
	makeSlice()
	nilSlice()
	appendEle()
}

func arr() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	// hello world [hello world]
	fmt.Println(a[0], a[1], a)

	var b [3]int
	b[2] = 4
	// [0 0 4]
	fmt.Println(b)
}

func slice() {
	fmt.Println("-------------")
	p := []int{4, 6, 7, 0, 9, 5}
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] => %d\n", i, p[i])
	}

	fmt.Println("p[:2]", p[:2])   // [4 6]
	fmt.Println("p[3:]", p[3:])   // [0 9 5]
	fmt.Println("p[1:4]", p[1:4]) // [6 7 0]
	fmt.Println("p[1:4]", p[3:3]) // []
}

func makeSlice() {
	a := make([]int, 4)
	// a: len=4 cap=4 [0 0 0 0]
	printSlice("a:", a)
	b := make([]int, 0, 5)
	// b: len=0 cap=5 []
	printSlice("b:", b)
	c := b[:3]
	// c: len=3 cap=5 [0 0 0]
	printSlice("c:", c)
	d := c[2:5]
	// d: len=3 cap=3 [0 0 0]
	printSlice("d:", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func nilSlice() {
	var a []int
	fmt.Println(a, len(a), cap(a)) // [] 0 0
	fmt.Println(a == nil)          // true

}

func appendEle() {
	var m []int
	// m: len=0 cap=0 []
	printSlice("m:", m)

	m = append(m, 3)
	// m: len=1 cap=1 [3]
	printSlice("m:", m)

	m = append(m, 0)
	// m: len=2 cap=2 [3 0]
	printSlice("m:", m)
	m = append(m, 7)
	// m len=3 cap=4 [3 0 7]
	printSlice("m:", m)
	m = append(m, 7)
	// m: len=4 cap=4 [3 0 7 7]
	printSlice("m:", m)
	m = append(m, 9)
	// m: len=5 cap=8 [3 0 7 7 9]
	printSlice("m:", m)
	m = append(m, 8, 1, 4, 2, 10, 45)
	// m: len=11 cap=16 [3 0 7 7 9 8 1 4 2 10 45]
	printSlice("m:", m)

	// cap: 1 -> 2 -> 4 -> 8 -> 16 -> ...
}

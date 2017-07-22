// var <varName> [n]<type>, n >= 0
package main

import "fmt"

func main() {
	var a [2]int
	var b [2]int
	b = a
	// [0 0]
	fmt.Println(b)

	var c [1]string
	// []
	fmt.Println(c)

	d := [3]int{1, 3, 5}
	// [1 3 5]
	fmt.Println(d)

	e := [10]int{9: 9}
	// [0 0 0 0 0 0 0 0 0 9]
	fmt.Println(e)

	f := [...]int{9: 9}
	// var p *[10]int = &f
	p := &f
	// [0 0 0 0 0 0 0 0 0 9]
	fmt.Println(f)
	// &[0 0 0 0 0 0 0 0 0 9]
	fmt.Println(p)

	// pointers array
	x, y := 5, 8
	g := [...]*int{&x, &y}
	// [0xc420072220 0xc420072228]
	fmt.Println(g)

	h := [2]int{1, 2}
	i := [2]int{1, 3}
	// true
	fmt.Println(h != i)

	j := new([5]int)
	j[2] = 3
	// &[0 0 3 0 0]
	fmt.Println(j)
	k := [5]int{}
	k[2] = 5
	// [0 0 5 0 0]
	fmt.Println(k)

	l := [2][4]int{
		{3: 4},
		{2, 4, 3: 4}}
	// [[0 0 0 4] [2 4 0 4]]
	fmt.Println(l)
}

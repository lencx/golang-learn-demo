package main

import "fmt"

func main() {
	var m map[int]string
	// m = make(map[int]string)
	m = map[int]string{}
	// OR
	// var m map[int]string = make(map[int]string)
	// OR
	// m := make(map[int]string)
	// map[]
	fmt.Println(m)
	m[1] = "Hello"
	m1 := m[1]
	// map[1:Hello] Hello
	fmt.Println(m, m1)
	delete(m, 1)
	// map[]
	fmt.Println(m)

	fmt.Println("********************")

	var mm map[int]map[int]string
	mm = make(map[int]map[int]string)
	a, ok := mm[1][1]
	if !ok {
		mm[1] = make(map[int]string)
		mm[2] = make(map[int]string)
	}
	mm[1][1] = "World"
	mm[1][2] = "HHHH"
	mm[2][1] = "AAAA"
	a, ok = mm[1][1]
	// World true
	fmt.Println(a, ok)
	// map[1:map[1:World 2:HHHH] 2:map[1:AAAA]]
	fmt.Println(mm)
}

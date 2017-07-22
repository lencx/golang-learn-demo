package main

import "fmt"

func main() {
	// array
	arr := [...]int{1, 3, 5, 7, 9}
	for i, v := range arr {
		// 0 1
		// 1 3
		// 2 5
		// 3 7
		// 4 9
		fmt.Println(i, v)
	}

	// map
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	for k, v := range m {
		// 1 A
		// 2 B
		// 3 C
		fmt.Println(k, v)
	}

	// slcie map
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "Hello"
		// map[1:Hello]
		// map[1:Hello]
		// map[1:Hello]
		// map[1:Hello]
		// map[1:Hello]
		fmt.Println(v)
	}
	// [map[] map[] map[] map[] map[]]
	fmt.Println(sm)

	for k := range sm {
		sm[k] = make(map[int]string, 1)
		sm[k][1] = "abc"
		// map[1:abc]
		// map[1:abc]
		// map[1:abc]
		// map[1:abc]
		// map[1:abc]
		fmt.Println(sm[k])
	}
	// [map[1:abc] map[1:abc] map[1:abc] map[1:abc] map[1:abc]]
	fmt.Println(sm)
}

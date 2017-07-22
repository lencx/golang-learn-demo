package main

import (
	"fmt"
)

// TODO:
// map[1:a 2:b 3:c 4:d 5:e]
// => map[a:1 b:2 c:3 d:4 e:5]

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	// map[4:d 1:a 2:b 3:c]
	fmt.Println(m)

	m2 := make(map[string]int)
	for k, v := range m {
		m2[v] = k
	}
	// map[c:3 d:4 a:1 b:2]
	fmt.Println(m2)
}

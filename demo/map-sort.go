package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E"}
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	// [3 4 5 1 2]
	// [5 1 2 3 4]
	// [3 4 5 1 2]
	fmt.Println(s)

	sort.Ints(s)
	// [1 2 3 4 5]
	fmt.Println(s)
}

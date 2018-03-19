package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Answer"] = 30
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 40
	fmt.Println("The value:", m["Answer"])

	fmt.Println(m)
	delete(m, "Answer")
	// map[]
	fmt.Println(m)

	v, ok := m["Answer"]
	// The value: 0 ; Present? false
	fmt.Println("The value:", v, "; Present?", ok)
}

package main

import "fmt"

func main() {
	i, j := 10, 20
	p := &i
	fmt.Println(&i, &j) // 0xc42000e368 0xc42000e390
	fmt.Println(*p)     // 10
	*p = 40
	fmt.Println(i) // 40

	p = &j
	*p = *p / 3
	fmt.Println(p, j) // 0xc42000e390 6
}

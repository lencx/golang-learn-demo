package main

import (
	"fmt"
)

type Add int

func main() {
	var n Add
	n.Increase(200)
	fmt.Println(n)
}

func (n *Add) Increase(num int) {
	*n += Add(num)
}

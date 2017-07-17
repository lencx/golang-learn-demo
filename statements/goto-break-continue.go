package main

import (
	"fmt"
)

func main() {
LABLE:
	for {
		for i := 1; i < 5; i++ {
			// 1
			fmt.Println("i is", i)
			break LABLE
		}
	}

	for {
		for j := 1; j < 10; j++ {
			// 1 - 6
			fmt.Println("j is", j)
			if j > 5 {
				goto LABLE2
			}
		}
	}
LABLE2:

LABLE3:
	for k := 1; k < 5; k++ {
		for {
			// 1 - 4
			fmt.Println("k is", k)
			continue LABLE3
		}
	}

LABLE4:
	for n := 0; n < 10; n++ {
		for {
			fmt.Println("n is", n)
			// 0 - 9
			continue LABLE4

			// `continue` replaced with `goto`
			// 0(infinite loop)
			// goto LABLE4
		}
	}
}

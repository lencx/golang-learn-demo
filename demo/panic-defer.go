package main

import (
	"fmt"
)

func main() {
	fn1()
	fn2()
	fn3()
}

func fn1() {
	fmt.Println("Func fn1")
}
func fn2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in fn2")
		}
	}()
	panic("Func fn2")
}
func fn3() {
	fmt.Println("Func fn3")
}

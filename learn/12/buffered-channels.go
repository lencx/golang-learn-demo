package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 3
	c <- 10
	// c <- 5
	fmt.Println(<-c)
	fmt.Println(<-c)

}

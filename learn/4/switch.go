package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("Hello")
	getOS()
	getWeek()
	getTime()
	deferFn()
}

func getOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s.", os)
	}
}

func getWeek() {
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func getTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferFn() {
	fmt.Println("counting")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}

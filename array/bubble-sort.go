package main

import "fmt"

func main() {
	arr := [...]int{2, 7, 5, 9, 1, 3, 6}
	// [2 7 5 9 1 3 6]
	fmt.Println(arr)
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	// [1 2 3 5 6 7 9]
	fmt.Println(arr)
}

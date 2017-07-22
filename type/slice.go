package main

import "fmt"

func main() {
	var s []int
	// []
	fmt.Println(s)

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(a)
	s1 := a[5]
	// 5
	fmt.Println(s1)

	// s2 := a[5:len(a)]
	// OR
	s2 := a[5:]
	// [5 6 7 8 9]
	fmt.Println(s2)
	s3 := a[:4]
	// [0 1 2 3]
	fmt.Println(s3)

	/* make([]T, length[, capacity]) */
	s4 := make([]int, 3, 10)
	// 3 10
	fmt.Println(len(s4), cap(s4))
	// [0 0 0]
	fmt.Println(s4)
	s5 := make([]int, 3)
	// 3 3
	fmt.Println(len(s5), cap(s5))

	arr := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	sArr := arr[1:5]
	// 4 5
	fmt.Println(len(sArr), cap(sArr))
	// bcde
	fmt.Println(string(sArr))

	fmt.Println("*******************************")
	/* reslice */
	ss := sArr[2:]
	// de
	fmt.Println(string(ss))

	/* append */
	s6 := make([]int, 3, 8)
	// [0 0 0]
	// 0xc420012200
	fmt.Printf("%p\n", s6)
	s7 := append(s6, 4, 5, 6)
	// [0 0 0 4 5 6] 0xc420012200
	fmt.Printf("%v %p\n", s7, s7)
	s8 := append(s7, 7, 8, 9, 10)
	// [0 0 0 4 5 6 7 8 9 10] 0xc420088000
	fmt.Printf("%v %p\n", s8, s8)
	// 16
	fmt.Println(cap(s8))

	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	sArr2 := arr2[3:5]
	sArr3 := arr2[1:4]
	// [4 5] [2 3 4]
	fmt.Println(sArr2, sArr3)
	sArr2[0] = 99
	// [99 5] [2 3 99]
	fmt.Println(sArr2, sArr3)
	// [1 2 3 99 5 6]
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	sArr4 := arr3[3:5]
	sArr5 := arr3[1:4]
	// [4 5] [2 3 4]
	fmt.Println(sArr4, sArr5)
	sArr5 = append(sArr5, 10, 11, 12, 13, 14, 15)
	sArr4[0] = 99
	// [99 5] [2 3 4 10 11 12 13 14 15]
	fmt.Println(sArr4, sArr5)

	/* copy */
	c1 := []int{1, 2, 3, 4, 5}
	c2 := []int{10, 11, 12}
	copy(c1, c2)
	// [10 11 12 4 5]
	fmt.Println(c1)

	c3 := []int{1, 2, 3, 4, 5}
	c0 := c3[:]
	// [1 2 3 4 5]
	fmt.Println(c0)
	c4 := []int{10, 11, 12}
	copy(c4, c3)
	// [1 2 3]
	fmt.Println(c4)
}

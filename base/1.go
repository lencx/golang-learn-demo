package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int = 5
	b := 3
	_, c := 7, 10
	fmt.Println(a, b, c)

	const Pi float32 = 3.14159
	const say = `hello,
Len`
	fmt.Println(Pi)
	fmt.Println(say)

	// bool
	var isActive bool
	isShow := false
	fmt.Println(isActive, isShow) // false false

	var c1 complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", c1)

	s := "hello"
	_s := "c" + s[1:]
	c2 := []byte(s)
	c2[0] = 'c'
	s2 := string(c2)
	fmt.Println(_s)
	fmt.Println(s2)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

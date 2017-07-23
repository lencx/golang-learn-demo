package main

import (
	"fmt"
)

type teacher struct {
	Name string
}

type student struct {
	Name string
}

type A int

func main() {
	t := teacher{Name: "Allen"}
	t.say()
	// {Mack}
	fmt.Println(t)
	s := student{Name: "Len"}
	s.say()
	// {Len}
	fmt.Println(s)

	var a A
	/* Method value */
	// Type is int
	a.say()
	/* Method expression */
	// Type is int
	(*A).say(&a)
}

func (t *teacher) say() {
	t.Name = "Mack"
	// Teacher is Mack
	fmt.Println("Teacher is", t.Name)
}

func (s student) say() {
	s.Name = "David"
	// Student is David
	fmt.Println("Student is", s.Name)
}

func (n A) say() {
	fmt.Println("Type is int")
}

package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type person2 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, Email, City string
	}
}

type human struct {
	Sex, Name string
	Age       int
}

type teacher struct {
	human
	Subject string
}

type student struct {
	human
	StudentID string
	Age       int
}

func main() {
	// { 0}
	fmt.Println(person{})
	// p1.Name = "Len"
	// p1.Age = 24
	p1 := person{"Len", 24}
	// {Len 24}
	fmt.Println(p1)

	/* value copy */
	// {Len 18}
	fn(p1)
	// {Len 24}
	fmt.Println(p1)

	/* address */
	// &{Len 18}
	pfn(&p1)
	// {Len 18}
	fmt.Println(p1)

	p2 := &person{"Joe", 16}
	// &{Joe 18}
	pfn(p2)
	p2.Name = "Tom"
	// &{Tom 18}
	fmt.Println(p2)

	/* anonymous struct */
	p3 := &struct {
		Name string
		Age  int
	}{
		Name: "L",
		Age:  20,
	}
	// &{L 20}
	fmt.Println(p3)

	p4 := person2{}
	// { 0 {  }}
	fmt.Println(p4)
	p5 := person2{Name: "Len", Age: 24}
	p5.Contact.Phone = "1234554321"
	p5.Contact.Email = "cxin1314@gmail.com"
	p5.Contact.City = "ShangHai"
	// {Len 24 {1234554321 cxin1314@gmail.com ShangHai}}
	fmt.Println(p5)

	var p6 person2
	p6 = p5
	// p6 is {Len 24 {1234554321 cxin1314@gmail.com ShangHai}}
	fmt.Println("p6 is", p6)

	t1 := teacher{}
	s1 := student{}
	// {{  0} } {{  0} }
	fmt.Println(t1, s1)
	t2 := teacher{human: human{Name: "Tony", Sex: "male", Age: 42}}
	s2 := student{human: human{Name: "Wendy", Sex: "female", Age: 22}}
	s2.Age = 25
	s2.human.Age = 28
	s2.human.Name = "Carey"
	// {{male Tony 42} } {{female Carey 28} 25}
	fmt.Println(t2, s2)
}

func fn(per person) {
	per.Age = 18
	fmt.Println(per)
}

func pfn(per *person) {
	per.Age = 18
	fmt.Println(per)
}

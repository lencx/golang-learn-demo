package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{2, 2}
	v2 = Vertex2{X: 7}
	v3 = Vertex2{}
	p  = &Vertex2{3, 5}
)

func main() {
	v := Vertex{4, 9}
	v.X = 8
	fmt.Println(v.Y) // 9
	fmt.Println(v)   // {8 9}

	p := &v
	p.X = 1e3
	fmt.Println(v)             // {1000 9}
	fmt.Println(v1, v2, v3, p) // {2 2} {7 0} {0 0} &{1000 9}
	fmt.Printf("%T\n", v1)     // main.Vertex2
	fmt.Printf("%T\n", p)      // *main.Vertex
}

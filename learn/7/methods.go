package main

import (
	"fmt"
	"math"
)

func main() {
	m1()
	m2()
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func m1() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func m2() {
	f := MyFloat(-math.Sqrt2)
	f2 := MyFloat(8)
	f3 := MyFloat(-8)
	fmt.Println(f.Abs(), f2.Abs(), f3.Abs())
}

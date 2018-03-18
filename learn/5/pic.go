package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		b := make([]uint8, dx)
		for y := range b {
			b[y] = uint8(x*y - 1)
		}
		a[x] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}

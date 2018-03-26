package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())        // (0,0)-(100,100)
	fmt.Println(m.At(0, 0).RGBA()) // 0 0 0 0
}

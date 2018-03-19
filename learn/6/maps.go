package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
	fmt.Println(m)
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// // map[Bell Labs:{40.68433 -74.39967}]
	// fmt.Println(m)

	// m["Google"] = Vertex{
	// 	37.42202, -122.08408,
	// }
	// // {37.42202 -122.08408}
	// fmt.Println(m["Google"])
}

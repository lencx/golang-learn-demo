package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[word]++
	}

	return wordCountMap
}

func main() {
	a := WordCount("hello world, hello")
	fmt.Println(a) // map[hello:2 world,:1]
}

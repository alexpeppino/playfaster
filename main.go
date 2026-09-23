package main

import (
	"fmt"
	"slices"
)

func main() {
	var sl = []int{1, 2, 4, 5}
	fmt.Printf("slice: %v\n", sl)
	sl = slices.Insert(sl, 2, 3)
	fmt.Printf("slice: %v\n", sl)
	sl = slices.Delete(sl, 1, 3)
	fmt.Printf("slice: %v\n", sl)
}

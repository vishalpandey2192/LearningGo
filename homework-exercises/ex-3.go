package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal; map literal
	x := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}
	for k, v := range x {
		fmt.Println(k, "-", v)
	}

	// struct
}

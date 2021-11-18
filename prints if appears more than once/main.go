package main

import (
	"fmt"
)

func main() {
	counts := make(map[string]int)
	var sliceOtString = []string{
		"Ivi",
		"Onyx",
		"Mishka",
		"Onyx",
		"Maslinka",
	}
	for _, elem := range sliceOtString {
		counts[elem]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	// slices
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// append item to slice
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// remove item from slice
	colors = append(colors[1:len(colors)]) // first item
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1]) // last item
	fmt.Println(colors)

	// make -- allows you to append
	numbers := make([]int, 5, 5) // array of ints, initial size 5, cap of 5
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 36
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	// sort array
	sort.Ints(numbers)
	fmt.Println(numbers)

}

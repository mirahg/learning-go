package main

import (
	"fmt"
)

func main() {
	// instantiate array and then assign members at index
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	// print all members of array
	fmt.Println(colors)
	// print member of array at index 0
	fmt.Println(colors[0])

	// instantiate array on same line
	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}

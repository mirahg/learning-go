package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// range -- more concise
	for i := range colors {
		fmt.Println(colors[i])
	}

	// for each loop
	for _, color := range colors {
		fmt.Println(color)
	}

	// no while keyword
	// use boolean condition instead
	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)

		// goto
		if sum > 200 {
			goto theEnd
		}
	}
	// what the goto calls
theEnd:
	fmt.Println("End of program")
}

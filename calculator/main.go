package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// define a reader object for user input
	reader := bufio.NewReader(os.Stdin)

	// ask for a value 1
	fmt.Printf("Value 1: ")
	// read the input for value 1 and check for proper input
	value1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		panic("Value 1 must be a number")
	}

	// ask for a value 2
	fmt.Printf("Value 2: ")
	// read the input for value 2
	value2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		panic("Value 2 must be a number")
	}

	sum := math.Round((float1+float2)*100) / 100
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", float1, float2, sum)
}

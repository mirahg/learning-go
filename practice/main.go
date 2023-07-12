package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// can also decalre the variable right after the 'switch' command
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's Sunday!"
		// fallthrough
	case 2:
		result = "It's Monday!"
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)
}

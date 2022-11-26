package main

import (
	"fmt"
)

func main() {
	// structs similar to java/C but no inheritance

	// create instance of struct
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	// change values
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// create custom type
// Dog is a struct
type Dog struct { // exported/public
	Breed  string
	Weight int
}

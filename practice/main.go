package main

import (
	"fmt"
	"sort"
)

func main() {
	// maps are like hash tables, unordered collection

	// create an empty map with key and value both strings
	states := make(map[string]string)
	fmt.Println(states)
	// add items to map
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	// access item by key
	california := states["CA"]
	fmt.Println(california)

	// delete item from map
	delete(states, "OR")
	// add new item
	states["NY"] = "New York"
	fmt.Println(states)

	// iterate through map
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// create a list of keys
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// sort the keys
	sort.Strings(keys)
	fmt.Println(keys)

	// print values, now ordered (sorted)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}

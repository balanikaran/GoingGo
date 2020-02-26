package main

import (
	"fmt"
)

// Set - Basic set implementation
// this can be improved using interfaces and recievers
type Set map[string]struct{}

func main() {

	// making an empty set
	s := make(Set)

	// adding values to set
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	// fetching back values from set
	fmt.Println(getValuesFromSet(s))
	fmt.Println(getValuesFromSet(s))
	fmt.Println(getValuesFromSet(s))
	// elements in map are not read in order, they are random, check the output

}

func getValuesFromSet(s Set) []string {
	var values []string
	for key, _ := range s {
		values = append(values, key)
	}
	return values
}
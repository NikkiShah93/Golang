package main

import (
	"fmt"
)

func main() {
	// we're going to go through maps in g in this module
	// now what are maps?
	// they're basically key-value pairs
	// one key type mapped to one value type
	// in the declaration
	// we can use a lot of types for key
	// and any type for value
	// but when maps are declared
	// the type has to be consistent
	// for every key-value pair within the map
	// the main constraint on the keys
	// is that it has to have the ability
	// to be tested for equality
	// so we can't use slices, maps and functions
	statePopulation := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
	}
	fmt.Println(statePopulation)
	// we can use arrays for keys however
	m := map[[2]int]string{
		{1, 2}: "one and two",
		{2, 3}: "two and three",
	}
	fmt.Printf("\n%v is of type %t\n", m, m)
}

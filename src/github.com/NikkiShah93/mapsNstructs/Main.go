package main

import (
	"fmt"
)

// when declaring a struct
// we have a collection of field names
// with their associated types
// and what it's doing is
// gathering info related
// to that concept
// we don't have any constraints
// on the type of data used for struct
// and we can mix any types
// which makes them very powerful
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

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
	// this is the most common way to create a map
	// map[key type]value type
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
	// the other way to create a map
	// we can also use the make function
	// and this is useful for cases
	// when you don't have the values
	// when creating the map
	// and want to populate it later
	new_map := make(map[string]int)
	new_map = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
	}
	fmt.Println(new_map)
	// now how can we manipulate our map
	// first way is to use [key] to access values
	fmt.Println(new_map["Texas"])
	// we can also add new values
	// by the following way
	new_map["Illinois"] = 12801539
	fmt.Println(new_map["Illinois"])
	// remeber that maps are not ordered
	// if we want ordered collections
	// arrays and slices would be suitable
	// we can also delete keys from map
	// using delete the following way
	delete(new_map, "Illinois")
	// if we interrogate a map
	// looking for a key that doesn't exist
	// it'll not throw an error
	// it'll return 0 for that key
	// which can be confusing
	// so if we want to distinguish
	// between the keys with value 0
	// and the missing ones
	// we can try to get info following way
	pop, ok := new_map["Illinois"]
	// and the second value is boolean
	// that shows if the key is present or not
	fmt.Println(pop, ok)
	// we can get the flag the following way too
	_, ok = new_map["Texas"]
	fmt.Println(ok)
	// we can use len func on maps also
	fmt.Println(len(new_map))
	// similar to slices
	// when you pass the map to a func
	// the original map will be impacted
	sp := statePopulation
	delete(sp, "Pennsylvania")
	// now the original map doesn't have this key either
	// so keep this in mind when passing maps around
	// because manipulating them can have side effects
	fmt.Println("statePopulation : ", statePopulation)
	fmt.Println("new map: ", sp)
	// now for struct
	// when we have at the package level
	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"Lisa Shaw"},
	}
	fmt.Println(aDoctor)
}

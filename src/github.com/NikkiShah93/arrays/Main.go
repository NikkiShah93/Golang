package main

import "fmt"

func main() {
	// arrays and their use cases
	// why do we need arrays?
	// when we need to store multiple related values
	// it can be cumbersome
	// so we use arrays instead
	// and this is how we declare them
	// we have to start with the size of the array
	// which is the number of elememts it's gonna hold
	// and the type of data it's gonna store
	// so it can only store one type of data
	// and we can initilaze the array
	// following way
	grades := [3]int{97, 85, 93}
	fmt.Printf("%v, %T\n", grades, grades)
	// another advantage of working with arrays
	// is the way they're laid out in the memory
	// because it's impossible to know
	// how different variable are laid out
	// by go runtime
	// but with arrays
	// we know they're contiguous in memory
	// meaning, accessing the various elemnets
	// of an array is very fast
	// if you're initilazing an array
	// while declaring it
	// you don't to specify the size
	// and could ask for an array
	// that's large enough to hold your data
	noSizeArray := [...]int{67, 87, 90}
	fmt.Printf("%v, %T\n", noSizeArray, noSizeArray)
	// we can also have empty arrays
	var students [3]string
	fmt.Printf("%v, %T\n", students, students)
	// and when assigning values to array
	// we need to specify the index
	students[0] = "Lisa"
	fmt.Printf("%v, %T\n", students, students)
}

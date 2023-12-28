package main

import (
	"fmt"
)

// our goal is to learn different datatypes
func main() {
	// first booleans
	// simple, similar to other langs
	var n bool = true
	fmt.Printf("value of n is %v, and type is %T\n", n, n)
	n = false
	fmt.Printf("value of n is %v, and type is %T\n", n, n)
	// using bools for logical testing
	f := 1 == 1
	m := 1 == 2
	fmt.Println(f, m)
	// in go
	// every time you initilaze a var
	// it has default value of zero
	// so for bools
	// it'll have value false
	var k bool
	fmt.Printf("%v, %T\n", k, k)
	// now for numeric types
	// staring with int
	// which is the default type for numbers
	// regardless of the env
	// an int will always be at least 32 bits
	// but it could be 32, or 64 too
	// depending on the system
	i := 42
	fmt.Printf("%v, %T", i, i)
}

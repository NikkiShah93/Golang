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
	// you can also have unsigned ints
	var ui uint16 = 42
	fmt.Printf("%v, %T\n", ui, ui)
	// we can't have uint of 64
	// but there's byte
	// which is the same as unit8
	// and that's used for data streaming
	// for encoding the data
	// you can't do arithmathic with two different types of ints
	// without converting one to the other
	// and the number generated
	// will always have the same type as the inputs
	a := 10 // in binary 1010
	b := 3  // in binary 0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	// we also have operators in go
	// that'll help you get the following
	// using the same variables
	fmt.Println("Operators in action:")
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	// either one has or the other does
	fmt.Println(a ^ b) // 1001
	// if neither one has it
	fmt.Println(a &^ b) // 0100
}

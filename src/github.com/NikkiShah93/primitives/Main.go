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
	// we also have bitwise operators in go
	// that'll help you get the following
	// using the same variables
	fmt.Println("Operators in action:")
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	// either one has or the other does
	fmt.Println(a ^ b) // 1001
	// if neither one has it
	fmt.Println(a &^ b) // 0100
	// the next thing is bit shifting
	w := 8 // 2^3
	// will shif left 3 places
	fmt.Println(w << 3) // 2 ^ 3 * 2 ^ 3
	// shift right 3 places
	fmt.Println(w >> 3) // 2 ^ 3 / 2 ^ 3
	fmt.Println(w >> 2) // 2 ^ 3 / 2 ^ 2
	fmt.Println(w << 2) // 2 ^ 3 * 2 ^ 2
	// now the floating point numbers
	// we can generate them the following way
	fl := 3.14   // if you let it infer the float
	fl = 13.7e72 // it can handle this
	// because it always use float64
	// otherwise, if you're using float 32
	// then it'll throw an error because
	// it can go to 10 to 38th power
	// because it'll be too big for it
	fl = 2.1e14
	fmt.Printf("%v, %T", fl, fl)
	// for arithmetic operations on floats
	// other than %, which is only available for ints
	// we also don't have bitwise or shifting operators
	// the rest will work as expected
	// returning a floating point number
	fn := 10.2
	fmt.Println(fl + fn)
	fmt.Println(fl - fn)
	fmt.Println(fl * fn)
	fmt.Println(fl / fn)
	// the next set is the complex numbers
	// which makes go a powerful language for data science
	// there are two types
	// 64 and 128
	// because we're using float32 | float64
	// for the real and imaginary parts
	var t complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", t, t)
	// and we all the operations that float has
	// for complex numbers as well
	// but then again
	// they should be the same type
	// similar to floats
	var y complex64 = 2 + 5.2i
	fmt.Println(t + y)
	fmt.Println(t - y)
	fmt.Println(t * y)
	fmt.Println(t / y)
	// you can aslo only access the real
	// or the imaginary parts of the number
	fmt.Printf("%v, %T\n", real(t), real(t))
	fmt.Printf("%v, %T\n", imag(t), imag(t))

}

package main

import "fmt"

// if you're defining a var outside main
// you have to declare the type
var t int = 8

func main() {
	// use this when you don't want
	// want to initilaze the var when defining it
	// and assign values later, maybe in a loop
	var i int
	i = 7
	// use this when go doesn't have enough info
	// to identify the type of data
	var j int = 99
	// use this when go has all the info
	// required to identify the type of data
	k := 42
	// you can't initilaze float32 using ^
	// because go will infer the numbers
	// to be int
	// or float64 if there's a decimal place
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	// use printf for formatted strings
	fmt.Printf("value is %v and type is %T", j, j)
}

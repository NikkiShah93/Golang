package main

import "fmt"

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
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}

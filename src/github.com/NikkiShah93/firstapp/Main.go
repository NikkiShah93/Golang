package main

import (
	"fmt"
	"strconv"
)

// if you're defining a var outside main
// you have to declare the type
var t int = 8

// other way you can define
// variables outside of the main block (at a packag level)
// is to warp them in a var
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)
var (
	counter int = 0
)

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
	fmt.Println()
	fmt.Println("T value is:", t)
	// if you declare a variable in the package level
	// and then you declare it again iside main
	// and try to declare it again
	// it'll throw an error
	// saying no new variable on the left side
	// but you can reassign a new value
	var counter int = 42
	// can't do this counter := 13
	// but can re-define inside the function
	// this is called shadowing
	counter = 13
	fmt.Println(counter)
	// one of the important things to remember
	// while using go
	// is that if you have un-used variables
	// it won't like it
	// and this will help keeping the application clean
	// this will help with improvements
	// becuase you won't have any old variables
	// sitting in your code
	// two sets of rules for naming your var
	// lowercase var are scoped to the package
	// so anything that works with it
	// can't see or use it
	// if it's uppercase
	// then go compiler
	// will expose this to other packages
	// and we also have the block scope
	// when you have var inside a func
	// it's never visible to ouside of the block itself
	// for naming var
	// if it's a simple var
	// that has a short lifespan
	// don't worry about it much
	// for other var
	// try longer more meaningful names
	// if you're working with acronym
	// use all uppercase
	var l int = 42
	fmt.Printf("value of l is %v, and the type is %T", l, l)
	// if we want to convert the type
	// we can do the following
	// you have to explicitly convert
	var m float32
	m = float32(l)
	fmt.Printf("the value of m is %v, and the type is %T", m, m)
	// but you can't do this with strings
	// because what go interpet the string to be
	// is to get the unicode char with that number
	var s string
	s = string(l)
	fmt.Println(s)
	// if you want to convert
	// number to strings back and forth
	// you need the strconv package
	// and do the following
	var ns string
	ns = strconv.Itoa(l)
	fmt.Println(ns)
}

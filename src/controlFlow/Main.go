package main

import (
	"fmt"
)

// this module is on different ways
// to control the flow of the app
// and the main topics are
// if , if-else, if-else-if
// and the switch statements
func main() {
	// lets start with if statements
	// we start with if keyword
	// then have a condition
	// that generates a boolean result
	// and inside the {} is the code
	// that'll be executed when
	// the condition is met
	// you can't have a signle line
	// blocks evaluating the result of an if statement
	if true {
		fmt.Println("The test is true")
	}
	// another common way is using
	// the initializer syntax
	statPopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	// which first gets some value
	// and then checks the boolean of those vals
	// notice there's ; between them
	// what comes before ; is part of
	// the initializer
	// these variables are defined within
	// the scope of the if statement
	if pop, ok := statPopulation["Florida"]; ok {
		fmt.Println(pop)
	}
	// now for the operators that can be used
	// for comparison of the condition
	// we have <, >, ==, <=, >=, !=
	// these work with the numeric types
	// but only != and == work with strings
	// we can also combine tests
	// with logical operators
	// with || (or) && (and)
	number := 50
	guess := 70
	if guess > 100 || guess < 1 {
		fmt.Println("the guess must be between 1 and 100")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("too low")
		}
		if guess > number {
			fmt.Println("too high")
		}
		if guess == number {
			fmt.Println("you got it!")
		}
	}

}

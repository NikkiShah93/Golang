package main

import (
	"fmt"
)

func main() {
	// all the constants
	// will start with the const keyword
	// the same way as it is with variables
	// if we want them to be exported
	// we need to have the first letter in uppercase
	// so this will be exported const MyConst
	// but this won't myConst
	// we can have typed and untyped constants
	// typed ones will be similar to variables
	const myConst int = 42
	fmt.Printf("%v, %T", myConst, myConst)
}

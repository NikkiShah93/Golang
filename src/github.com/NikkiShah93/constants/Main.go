package main

import (
	"fmt"
)

const f int16 = 17

// enumerated constants
const ec = iota

// we can also do the following
const (
	ec1 = iota
	ec2 = iota
	ec3 = iota
)

// or having the compiler infer the values
const (
	ec4 = iota
	ec5
	ec6
)

// or use the following approach
// you can also add numbers to iota
// and it won't throw an error
// so you can start from another number
// if your application requires
const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

// we can also use bit shifting on iota
// this way we can raise things
// to power of two
// shifting by one level
// you're multiplying it by two
const (
	_  = iota // ignoring the first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

// we can also use this for setting flags
// and have boolean flags in one single byte
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
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
	// and constants have to be assignable at compile time
	// which means you can't have your constants being determind at runtime
	// so this will work
	const myConst int = 42
	// but this won't
	// const myConst float64 = math.sin(1.57)
	// which needs to be calculated at runtime
	// because executing functions is not allowed at compile time
	// which includes the flags that you're passing into your app
	fmt.Printf("%v, %T\n", myConst, myConst)
	// constant can be made of any of the primitive types
	const a int = 2
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	// constants can also be shadowed
	// and inner definition of the variale always wins
	const f int16 = 27
	fmt.Printf("%v, %T\n", f, f)
	// they're also very similar to variable
	// when it comes to operations
	// and we can even use them in combination
	// with variables
	// as long as they're the same type
	var e int16 = 10
	fmt.Printf("%v, %T\n", f+e, f+e)
	// we can also have the compiler infer the type
	const g = 8
	fmt.Printf("%v, %T\n", g, g)
	// but when the compiler sees that the const
	// is used in combination with another type
	// it'll infer it to be of that othe type
	fmt.Printf("%v, %T\n", g+e, g+e)
	// next thing is enumerated constants
	// defined in package level
	// and is used as a counter
	fmt.Printf("%v, %T\n", ec, ec)
	// when defining them as a const block
	// then they'll increase by 1 each time
	// and it's scoped to one const block
	// meaning, it will reset in each block
	fmt.Printf("%v, %T\n", ec1, ec1)
	fmt.Printf("%v, %T\n", ec2, ec2)
	fmt.Printf("%v, %T\n", ec3, ec3)
	// using the const block in package level
	var specialistType int = catSpecialist
	// this will be true even if we don't initilaze the variable
	// because the default will be 0
	// which is the same as the first value of iota set
	// so it's better practice to use the 0th iota
	// to be the error const
	// so it won't lead to mistakes in your logic
	// or maybe assign the 0th value to _
	// so you won't have to assign memory to it
	fmt.Printf("%v\n", specialistType == catSpecialist)
	//using bit shifting with iota
	fileSize := 400000000.
	// formatting by 2 decimal places
	// for a floating number
	fmt.Printf("%.2fGB\n", fileSize/GB)
	// using the iota for flags
	var roles byte = isAdmin | isHeadquarters | canSeeFinancials
	fmt.Printf("%b\n", roles)
	// this is a very efficient way to store flags
	// because you can check is someone is an admin
	// by the follwing logic
	fmt.Printf("Is Admin? %v\n", roles&isAdmin == isAdmin)

}

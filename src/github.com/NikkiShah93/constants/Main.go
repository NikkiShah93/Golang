package main

import (
	"fmt"
)

const f int16 = 17

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
}

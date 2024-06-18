package main

import "fmt"

func main() {
	// arrays and their use cases
	// why do we need arrays?
	// when we need to store multiple related values
	// it can be cumbersome
	// so we use arrays instead
	// and this is how we declare them
	// we have to start with the size of the array
	// which is the number of elememts it's gonna hold
	// and the type of data it's gonna store
	// so it can only store one type of data
	// and we can initilaze the array
	// following way
	grades := [3]int{97, 85, 93}
	fmt.Printf("%v, %T\n", grades, grades)
	// another advantage of working with arrays
	// is the way they're laid out in the memory
	// because it's impossible to know
	// how different variable are laid out
	// by go runtime
	// but with arrays
	// we know they're contiguous in memory
	// meaning, accessing the various elemnets
	// of an array is very fast
	// if you're initilazing an array
	// while declaring it
	// you don't to specify the size
	// and could ask for an array
	// that's large enough to hold your data
	noSizeArray := [...]int{67, 87, 90}
	fmt.Printf("%v, %T\n", noSizeArray, noSizeArray)
	// we can also have empty arrays
	var students [3]string
	fmt.Printf("%v, %T\n", students, students)
	// and when assigning values to array
	// we need to specify the index
	// and go knows the pointer to this string
	// and it'll grab the index element
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("%v, %T\n", students, students)
	// and we can use [] to get the element
	fmt.Printf("student #1 : %v\n", students[1])
	// we can also access the size of the array
	fmt.Printf("The size of the array is %v\n", len(students))
	// we can have arrays for any type
	// but within one array,
	// elements should be of same type
	// we can also have array of arrays
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// we can also initilaze each indivisually
	// identityMatrix[0] = [3]int{1,0,0}
	// and so on
	fmt.Println(identityMatrix)
	// in go, arrays are considered values
	// in other langs
	// the array is pointing to the values in that array
	// but in go, when you're creating an array
	// and making a copy
	// it'll generate a literal copy
	// which is not pointing to the same values
	// so you can re-assign the entire array
	// without changing the first version
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Printf("a is %v\nb is %v\n", a, b)
	// if you don't want to have this behavior
	// we can pass the pointer instead
	// by adding & to the var name
	c := [...]int{4, 5, 6}
	d := &c
	d[1] = 7
	fmt.Printf("c is %v\nd is %v\n", c, d)
	// the fact the size should be known at compile time
	// limits the array
	// and we can use slices instead
	// which are almost the same as arrays
	// in terms of funtionalities
	// with a few exceptions
	e := []int{1, 2, 3}
	fmt.Println(e)
	// we have both length and capacity for slices
	// which could be different
	fmt.Printf("the length of the slice is %v, and its capacity is %v\n", len(e), cap(e))
	// the other difference between arrays and slices
	// is that the copy of a slice, will point to the same values
	// and we don't need to use pointers for them
	f := e
	f[1] = 5
	fmt.Printf("e is %v\nf is %v\n", e, f)
	// there are several other ways to create slices
	// which are common between arrays and slices
	g := f[:]  // all elements of f
	h := f[1:] // from 2nd element
	i := f[:1] // copy till 1st
	// remeber, 1st number is inclusive
	// and the 2nd number is exclusive
	fmt.Printf("f is %v\ng is %v\nh is %v\ni is %v\n", f, g, h, i)
	// we can also use the make func
	// first we pass the type, then the length
	// and we can also pass the capacity (optional)
	j := make([]int, 3, 100)
	fmt.Printf("the length is %v and the capacity is %v\n", len(j), cap(j))
	// we can append to the slices
	j = append(j, 1)
	fmt.Printf("%v has the length of %v and the capacity of %v\n", j, len(j), cap(j))
	// using make and passing capacity is better practice
	// for cases where the slices can be large,
	// because each time we're reshaping a slice
	// go will create a new copy, so it can handle more elements, if needed
	// and it will double size of the needed array and add it to capacity
	// if 0 > add 1 > capacity > 2
	// if len 1 > add 3 > capacity > 8
	// for append function
	// everything after the 1st element
	// is cosidered a value to be appended
	j = append(j, 4, 5, 6, 7)
	fmt.Printf("%v has the length of %v and the capacity of %v\n", j, len(j), cap(j))
	// you can't append two slices together
	// this way append(a,b)
	// there's a way around it
	// which is spreading out the elements
	// by adding ... at the end
	j = append(j, []int{9, 10, 11}...)
	fmt.Printf("%v has the length of %v and the capacity of %v\n", j, len(j), cap(j))
	// another common operations with slices
	// is to treat them like stacks
	// and for poping the elements from a slice
	// we can only copy the elements we want to another slice
	// but be careful
	// because we are always dealing with the same underlying arrays
	// so any changes later on on the copy
	// will impact the original slice
	// and there's no way around this
}

package main

import (
	"fmt"
	"reflect"
)

// when declaring a struct
// we have a collection of field names
// with their associated types
// and what it's doing is
// gathering info related
// to that concept
// we don't have any constraints
// on the type of data used for struct
// and we can mix any types
// which makes them very powerful
// we can even have structs in structs
// the naming rules are the same for structs
// if it starts with capital letter
// it'll be exported from the package
// otherwise, it won't
// so if we want the fields to be exported
// with the main struct, we have to
// capitalize the names
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// the use of embedding example
type Animal struct {
	Name   string
	Origin string
}

// in order to create similar
// to inheritance relationships
// we just add one struct to another
// the following way
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// example for tags
// the values of tags can be arbitrary
// you could use any strings
// but the conventional use within go
// is to have space delimited sub tags
// and for key-value pairs
// use colon in between them
// and the values are in quotation marks
type User struct {
	Name string `required max:"100"`
	Age  int16  `required`
}

func main() {
	// we're going to go through maps in g in this module
	// now what are maps?
	// they're basically key-value pairs
	// one key type mapped to one value type
	// in the declaration
	// we can use a lot of types for key
	// and any type for value
	// but when maps are declared
	// the type has to be consistent
	// for every key-value pair within the map
	// the main constraint on the keys
	// is that it has to have the ability
	// to be tested for equality
	// so we can't use slices, maps and functions
	// this is the most common way to create a map
	// map[key type]value type
	statePopulation := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
	}
	fmt.Println(statePopulation)
	// we can use arrays for keys however
	m := map[[2]int]string{
		{1, 2}: "one and two",
		{2, 3}: "two and three",
	}
	fmt.Printf("\n%v is of type %t\n", m, m)
	// the other way to create a map
	// we can also use the make function
	// and this is useful for cases
	// when you don't have the values
	// when creating the map
	// and want to populate it later
	new_map := make(map[string]int)
	new_map = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
	}
	fmt.Println(new_map)
	// now how can we manipulate our map
	// first way is to use [key] to access values
	fmt.Println(new_map["Texas"])
	// we can also add new values
	// by the following way
	new_map["Illinois"] = 12801539
	fmt.Println(new_map["Illinois"])
	// remeber that maps are not ordered
	// if we want ordered collections
	// arrays and slices would be suitable
	// we can also delete keys from map
	// using delete the following way
	delete(new_map, "Illinois")
	// if we interrogate a map
	// looking for a key that doesn't exist
	// it'll not throw an error
	// it'll return 0 for that key
	// which can be confusing
	// so if we want to distinguish
	// between the keys with value 0
	// and the missing ones
	// we can try to get info following way
	pop, ok := new_map["Illinois"]
	// and the second value is boolean
	// that shows if the key is present or not
	fmt.Println(pop, ok)
	// we can get the flag the following way too
	_, ok = new_map["Texas"]
	fmt.Println(ok)
	// we can use len func on maps also
	fmt.Println(len(new_map))
	// similar to slices
	// when you pass the map to a func
	// the original map will be impacted
	sp := statePopulation
	delete(sp, "Pennsylvania")
	// now the original map doesn't have this key either
	// so keep this in mind when passing maps around
	// because manipulating them can have side effects
	fmt.Println("statePopulation : ", statePopulation)
	fmt.Println("new map: ", sp)
	// now for struct
	// when we have at the package level
	// and this is how we create a struct
	// using the {} like any other literal definitions
	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"Lisa Shaw", "Jo Grant", "Sarah Jane Smith"},
	}
	fmt.Println(aDoctor)
	// in order to access a value from a struct
	// we can use the dot syntax
	// the following way
	fmt.Println("The actor name is", aDoctor.actorName)
	// we can even drill down through the structure
	fmt.Println(aDoctor.companions[0])
	// the field names are not required when initiating a struct
	// but this way is not recommended
	// because it'll cause problems if the original structure
	// is modified, because of the positional nature
	// will require the correct order and number of fields
	newDoctor := Doctor{
		4,
		"John Doe",
		[]string{"Jane Doe"},
	}
	fmt.Println(newDoctor)
	// we can declare the structs the following way as well
	// which will be an anonymous struct
	// which can't be used anywhere else
	// because it doesn't have an independent name
	// that we can refer to later on
	// the first set of {} is defining the structure
	// the second {} is the initializer
	// this isn't used that often
	// and only for short lived structs
	aCity := struct{ name string }{name: "Toronto"}
	fmt.Println(aCity)
	// unlike maps, structs are value types
	// meaning, if we copy and maipulate the data
	// the original struct will not be impacted
	// because it passes a copy around
	// so if you have a large struct
	// keep in mind that each round of manipulation
	// will create a new copy of that struct
	// similar to arrays, we can change this behavior
	// by passing the pointer to the struct
	anotherCity := aCity
	anotherCity.name = "Dallas"
	fmt.Println(aCity)
	fmt.Println(anotherCity)
	yetAnotherCity := &anotherCity
	yetAnotherCity.name = "Los Angeles"
	fmt.Println(anotherCity)
	fmt.Println(yetAnotherCity)
	// now the embedding concept
	// go language doesn't support
	// traditional oo principles
	// we don't have inheritance
	// but go allows for embedding
	// which could be used instead
	// it uses a model similar to
	// inheritance, called composition
	// and it's supported through embedding
	aBird := Bird{}
	aBird.Name = "Emu"
	aBird.Origin = "Australia"
	aBird.SpeedKPH = 48
	aBird.CanFly = false
	fmt.Println(aBird)
	// we can even dig into the fields
	// that are coming from the Animal struct
	fmt.Println("The bird's name is", aBird.Name)
	// but bird and animal are still distinct types
	// so despite the fact that we have Animal fields
	// Bird is not of Animal type
	// but a Bird has an Animal
	// if we want to use them interchangeably
	// we have to use interfaces
	// if we want to initilaze the Bird
	// using the literal syntax
	// we need to know the structure
	// and explicitly use the Animal struct
	anotherBird := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(anotherBird)
	// embedding is not the right choice for modeling
	// it's only useful when you have a sophisticated struct
	// that has certain functionalities
	// that can be useful in other structs
	// the next concept for structs are tags
	// which is useful if you're dealing with web app
	// and want certain limits/constraints on fields
	// and this is how we can access the tags
	// and then something else has to actually use these
	// because they're basically just strings
	t := reflect.TypeOf(User{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

package main

import (
	"fmt"
)

func main() {

	/*
		Structs are like maps, but is a less restrictive version of maps.
		Structs are also key-value pairs, but the keys can be of different data
		types and values can be different data types too.

		However the fields are fixed at compile type. I.e. you can't add/remove key/value pair
		from a struct during run time. e.g. you can add a new 'gender' field in
		the 'employee' example struct given below.

		For that reason, structs are similar to classes,
		in the sense that creating a struct entity is a bit like creating an object.

		It's also a 2-step-process to use structs. You first have to 'define' a struct,
		then you can 'instantiate' objects from the struct. You define structs using
		the 'type' key word:
	*/
	type employee struct {
		ID        int
		Title     string
		FirstName string
		Lastname  string
		Age       int
		Smoker    bool
	}

	/*
	  Now you can define variable as the (struct) datatype of 'employee'
	*/
	var david employee
	fmt.Println(david) // {0    0 false}
	// i.e. default string and int values

	// To print a particular field we use the . operator, to drill down
	fmt.Println(david.ID) // 0
	// here's how to set/override a field
	david.ID = 10
	fmt.Println(david)    // prints {10    0 false}
	fmt.Println(david.ID) // prints 10

	// Here's the more common way to initailise+declare a struct datatype
	Charlie := employee{
		ID:        23,
		Title:     "Mr",
		FirstName: "Charles",
		Lastname:  "Dickens",
		Age:       55,
		Smoker:    true,
	}

	fmt.Println(Charlie) // {23 Mr Charles Dickens 55 true}

	// Here's a more shorthand way to write a struct
	john := employee{34, "Dr", "John", "Smith", 30, false}
	fmt.Println(john) // {34 Dr John Smith 30 false}
	// However this approach is bad practice because it is harder to read the code.

	// Theres two syntax to construct an array of structs. Here I'm showing boths syntax in a combined example:
	employees := []employee{{
		ID:        23,
		Title:     "Mr",
		FirstName: "Charles",
		Lastname:  "Dickens",
		Age:       55,
		Smoker:    true,
	}, david, john}
	fmt.Println(employees)

}

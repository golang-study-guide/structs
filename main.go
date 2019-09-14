package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		Maps are a collection of key-value pairs, where all the keys have to be of a certain
		datatype, and all the values have to belong to a specific data type
		Here we are creating a map, where they keys are all strings, and the value are all
		integers
	*/
	weekDays := map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Sunday":    5,
	}
	fmt.Println(reflect.TypeOf(weekDays)) // prints map[string]int
	fmt.Println(weekDays)                 // prints map[Monday:1 Sunday:5 Thursday:4 Tuesday:2 Wednesday:3]
	// notice that it's in alphabetical order
	fmt.Println(weekDays["Wednesday"]) // prints 3

	// Here's how to replace a value:
	weekDays["Wednesday"] = 25
	fmt.Println(weekDays) // prints map[Monday:1 Sunday:5 Thursday:4 Tuesday:2 Wednesday:25]

	// Here's how to add a new key-value pair
	weekDays["Saturday"] = 6
	fmt.Println(weekDays) // prints map[Monday:1 Saturday:6 Sunday:5 Thursday:4 Tuesday:2 Wednesday:25]

	// Here's how to delete a key-value pair
	delete(weekDays, "Monday")
	fmt.Println(weekDays) // prints map[Saturday:6 Sunday:5 Thursday:4 Tuesday:2 Wednesday:25]

}

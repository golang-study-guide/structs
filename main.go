package main

func main() {

	/*
		Structs are like maps, but is a less restrictive version of maps.
		Structs are also key-value pairs, but the keys can be of different data
		types and values can be different data types too.

		However the fields are fixed at compile type. I.e. you can't add/remove key/value pair
		from a struct during run time. For that reason, structs are similar to classes,
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
	 */
}

package data

import "fmt"

// TYPE ALIASES

type number = int16

var edad number = 16

// Is possible to create aliases for types, like in this example
// every new value of type number would be like an int16 type
// For defining aliases or new types, the keyword type is used.

// NEW TYPES AND METHODS
type distance float32
type location string

// We can create new types using the type keyword and asign to that
// type a base like string in the example. Later that type can be
// enhanced via methods.

func (origin location) DistanceTo(destination location) distance {
	//TODO calculations...
	fmt.Printf("Origin %v Destination %v", origin, destination)
	return 10
}

//this function contains a special argument before its name that
//indicates that this function is a method of the type used in that
//argument. This method is enhancing location and can be called using
//the dot notation over a location type variable.

func locationTest() {
	nyc := location("33.4234, 34.456")

	nyc.DistanceTo(location("-23, -44"))
	// In this example we can see how a variable of type location can
	// call the method to calculate the distance to another location
	// and return a distance value.

	print(nyc)
}

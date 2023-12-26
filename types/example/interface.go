package data

type Signable interface {
	Signup() bool
}

//Interfaces are special types to group different types under a same
//name. To use an interface the keyword interface is used, and inside
//the curly brackets is necessary to include a list of methods
// are common to all the types we want group.
// the method list should only include the name signature and
// return type.

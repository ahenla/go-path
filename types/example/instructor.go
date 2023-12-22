package data

import "fmt"

//STRUCTURES

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// structures are a new type of value with definite properties
// that can have different types each one of them.
// this structures needs to be defined starting with the keyword
// type, followed by the name of the structure and the struct keyword
// after that and the properties inside curly braces.

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", i.LastName, i.FirstName, i.Score)
}

//this example shows the way to add a method to the Instructor struct
//to be used elsewhere you want to call it.

func NewInstructor(firstname string, lastname string) Instructor {
	return Instructor{
		FirstName: firstname,
		LastName:  lastname}
}

//This is a design pattern known as a structure factory. A function that
//return a struct with the arguments as properties of the structure.

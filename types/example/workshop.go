package data

import "time"

//EMBEDDING TYPES

type Workshop struct {
	Course
	Instructor
	Date time.Time
}

//It is possible to embed a struct inside another using the syntax from
//the example. Only using the name of the struct you want to embed
//as the name property without the type indicates that it is an
//embedded struct.
//The new Struct will inherit all the properties and methods of the
//embedded structures.
//In case of name collision is possible to use the dot sintax to call
//an specific property of an embedded structure.

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}

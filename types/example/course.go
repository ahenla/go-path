package data

import "fmt"

type Duration float32 //in hours

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) String() string {
	return fmt.Sprintf("---%v ---%v", c.Name, c.Instructor.FirstName)
}

// this method uses a special syntax with the name String that returns
// a string to modify the formatting output of the printed structure.

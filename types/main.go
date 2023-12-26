package main

import (
	"fmt"

	data "fe.com/go/types/example"
)

func main() {
	max := data.Instructor{
		Id:        1,
		FirstName: "Maximiliano",
		LastName:  "Firtman",
	}
	//new Instructor structure called.
	max.Score = 10

	kyle := data.NewInstructor("Kyle", "Simpson")

	goCourse := data.Course{
		Id:         2,
		Name:       "Go fundamentals",
		Instructor: max,
	}

	swiftWS := data.NewWorkshop("Swift with iOs", max)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	print(max.Print())
	print("\n")
	print(kyle.Print())
	print("\n")

	fmt.Printf("%v", goCourse)
	print("\n")
	fmt.Printf("%v", swiftWS)
	print("\n")

	for _, course := range courses {
		fmt.Println(course)
	}

}

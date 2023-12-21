package main

import data "fe.com/go/types/example"

func main() {
	max := data.Instructor{
		Id:        1,
		FirstName: "Maximiliano",
		LastName:  "Firtman",
	}
	//new Instructor structure called.

	max.Score = 10

}

package data

import "fmt"

// curiosities of function execution


func init (){

	var age int

	defer fmt.Println("Bye!")
	defer fmt.Println("Good")

	/*
	defer is a special keyword that defers the code after that line
	to the end of the execution of a function.
	In the case of more than one defer in the same function, the
	firs declared is the last executed
	*/

	if(age > 39) {
		panic("too old to be true")
	}
	/*
	panic() is a special function that will interrupt your program
	execution and display a message.
	in case a defer line is defined in that program, it will be
	executed before closing the app
	*/
}

package data //the package use the name of the folder is in

import "fmt"

var Countries [10]string // array definition, is static


//if you start the name ofa variable with a capital letter
// it means that that variable can be exported

var Slice []int //slices are like arrays but the size is dynamic

var Codes map[string]int //maps are collections of key value pairs
//map is mandatory, followed for the type of the key between brackets
//and the type of the value after the brackets

func init(){ //init is a special function that runs first and
	//can be used to initialized things before main.
	// it is possible to have multiple init functions in the same file

	Countries[0] = "Chile"
	Countries[3] = "Brazil"
	Countries[9] = "Japan"

	length := len(Countries)

	fmt.Println(length, " countries")
}

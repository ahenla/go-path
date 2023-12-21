package main

import "fmt"

var url = "https://andreshenriquez.me" // global variable

func main(){
	var message string = "hello from go" //explicit type
	var price = 34.4 //implicit type
	currency := "dollars" //shortcut declaration/initialization

	fmt.Print(message, " ")
	fmt.Println(price, currency)
}

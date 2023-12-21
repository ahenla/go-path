package main

import "fmt"

func main(){
	var operation string
	var number1, number2 int


	fmt.Println("  CALCULATOR GO 1.0")
	fmt.Println("=======================")

	fmt.Println("\nWhich operation you want to perform?")
	fmt.Println("add, substract, multiply, divide")
	fmt.Scanf("%s", &operation)

	fmt.Println("\nEnter first number")
	fmt.Scanf("%d", &number1)

	fmt.Println("\nEnter second number")
	fmt.Scanf("%d", &number2)

	fmt.Println("\nThe result is:")

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "substract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	default:
		println("*****invalid operation****")

	}
}

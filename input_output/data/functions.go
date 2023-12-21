package data

import "fmt"

func CalculateTax(price float32) (float32, float32){
	return price * 0.09, price * 0.02
}

/*
functions in go pass arguments by value, the parameters are
declared with the respective type.
the type of the return value, if there is one, is declared before
the curly brackets.
functions in go can return more than one value at the same time
you need to receive those two values with two variables or use underscore
_ in the place of the value you want to ignore.
*/

func CalculateTaxNamed(price float32) (stateTax float32, cityTax float32){
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}

/* It is possible to name the returned values for internal use inside
the function*/

func init(){
	stateTax, cityTax := CalculateTax(100)
	countryTax, _ := CalculateTax(50)

	fmt.Println(countryTax, stateTax, cityTax)
}

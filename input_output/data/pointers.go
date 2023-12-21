package data

import "fmt"

func Birthday(age int){
	age++
}
/*functions receive arguments by the value, like a copy,
so they can not modify the actual value of a variable*/

func PointerBirthday (pointerAge *int){
	*pointerAge++
}
/* for modifying the actual value of variables is necessary
pass a pointer to the value using the * sign before defining
the type of the parameter
	For accessing the value of the pointer inside the function
the * sing is also used before the argument name*/

func init(){
	age := 22
	Birthday(age)
	fmt.Println(age)
}
//this function won't modify the age variable

func init(){
	age := 22
	PointerBirthday(&age)
	fmt.Println(age)
}
/* when calling a function that receives a pointer is necessary
to use the & sign before the argument name to indicate you are
referring to the variable's pointer*/

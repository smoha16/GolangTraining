package main

import "fmt"

var y = 23

//DECLARED the VARIABLE with the IDENTIFIER "z"
//is of TYPE string
//cannot be changed inside the func main() of any other type(like int/bool/float)
//STATIC Programming Language
//Not a DYNAMIC Programming Language
//A VARIABLE is DECLARED to hold a VALUE of a certain TYPE
var z = "Shaken,not stirred"
var a string = `Hello brother`

//"y" is default type "int"
// "%T"=shows type	"\n"=newline	here we use fmt.Printf
func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	z = `hi there`
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}

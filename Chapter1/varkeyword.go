package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z bool

func main() {
	x := 23
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

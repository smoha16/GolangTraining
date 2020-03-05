package main

import "fmt"

var y string
var z int

func main() {

	fmt.Println("hello", y, "are you there")
	fmt.Printf("%T\n", y)

	y = `my name is bond`
	fmt.Println("hello", y, "are you there")
	fmt.Printf("%T\n", y)

	fmt.Println("hello", z, "are you there")
	fmt.Printf("%T", z)
}

package main

import "fmt"

type shanky int

var a shanky
var b int

func main() {
	a = 22
	b = 23
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = shanky(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

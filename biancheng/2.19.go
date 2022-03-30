package main

import "fmt"

type NewInt int

type IntAlias = int

func main() {
	var a NewInt
	var b IntAlias

	a = 1
	b = 1
	fmt.Printf("a type is %T\n", a)
	fmt.Printf("b type is %T\n", b)
}
package main

import (
	"fmt"
)

const a = 42

func main() {
	const a = 12
	/*
		you cannot assign a const: i = 15 is illegal
	*/
	fmt.Printf("%v, %T \n", a, a)
	fmt.Println("Hello Playground")
}

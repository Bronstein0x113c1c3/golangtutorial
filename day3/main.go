package main

import (
	"fmt"
)

func main() {
	const i = 42
	/*
		you cannot assign a const: i = 15 is illegal
	*/
	fmt.Printf("%v, %T \n", i, i)
	fmt.Println("Hello Playground")
}

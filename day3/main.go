package main

import (
	"fmt"
)

const (
	red = iota
	yellow
	blue
	black
)

func main() {
	// const a = 12
	/*
		you cannot assign a const: i = 15 is illegal
	*/
	/*
		inner-scope variable is more prioritized
	*/
	// fmt.Printf("%v, %T \n", a, a)
	// fmt.Println("Hello Playground")

	fmt.Printf("%v, %T \n", red, red)
	fmt.Printf("%v, %T \n", yellow, yellow)
	fmt.Printf("%v, %T \n", blue, blue)
	fmt.Printf("%v, %T \n", black, black)

	/*
		iota start from zero, in the list of constants
	*/

}

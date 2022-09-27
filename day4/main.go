package main

import "fmt"

func main() {
	a := [3]int{2, 5, 10}
	// points[0] =
	// points[1] =
	// points[2] =
	// b := a  // copy an array
	b := &a //copy the address of a to b
	b[1] = 20
	fmt.Println(a)
	fmt.Println(b)
}

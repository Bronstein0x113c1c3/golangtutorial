package main

import "fmt"

func main() {
	a := make([]int, 10)
	// points[0] =
	// points[1] =
	// points[2] =
	// b := a  // copy an array
	// b := a[3:6]
	// b:= a[x:y] copy from a from index x to index y-1
	a = append(a, 1)
	fmt.Printf("a: %v %v %v", a, len(a), cap(a))
	a = append(a, 2)
	fmt.Printf("a: %v %v %v", a, len(a), cap(a))
}

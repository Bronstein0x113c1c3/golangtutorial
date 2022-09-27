package main

import "fmt"

func main() {
	// a := []int{2, 5, 10, 12, 45, 50}
	a := make([]int, 10)
	// points[0] =
	// points[1] =
	// points[2] =
	// b := a  // copy an array
	b := a[3:6]
	// b:= a[x:y] copy from a from index x to index y-1
	fmt.Println(b)
}

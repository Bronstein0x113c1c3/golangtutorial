package main

import (
	set "ex6/intset"
	"fmt"
)

func main() {
	var x set.IntSet
	x.Add(1)
	x.Add(242)
	x.Add(144)
	x.Add(235)
	fmt.Println(&x)
	var y set.IntSet
	y.Add(242)
	y.Add(180)
	y.Add(235)
	y.Add(75)
	z := x.DifferenceWith(&y)
	z = x.SymmetricDifference(&y)
	fmt.Println(z)
	fmt.Println(x.Elems())
}

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
	fmt.Println(&x)

}

package main

import (
	"fmt"
	"test_algo/processing"
)

func main() {
	input := "50000*(3+12)*423*(479+21) "
	fmt.Println(processing.MainProcess(input))
}

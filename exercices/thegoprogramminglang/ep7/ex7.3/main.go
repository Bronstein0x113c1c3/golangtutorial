package main

import (
	"math/rand"

	"ex7/treesort"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.PrintValues(data)
}

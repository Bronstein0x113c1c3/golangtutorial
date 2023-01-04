package main

import "fmt"

func pascal_triangle_concurrency(x []int) []int {
	var b, c int
	a := x

	number_chan := make(chan int, 2)
	result_chan := make(chan int, 1000000)
	go func() {
		for {
			b = <-number_chan
			c = <-number_chan
			result_chan <- b + c
		}
	}()
	for i := 1; i <= len(a)-1; i++ {
		number_chan <- a[i-1]
		number_chan <- a[i]
		if i >= 3 {
			a[i-2] = <-result_chan
		}

	}
	return a
}

func main() {

	/*if the length of arr is x, Pascal Triangle could reach x-2 units
	, with the lowest unit start from index 0 => highest index is x-3!
	*/
	a := make([]int, 18)
	a[0] = 1
	a[1] = 0

	fmt.Println(a[0])
	for i := 1; i <= len(a)-3; i++ {
		a = pascal_triangle_concurrency(a)
		for j := 0; j <= i; j++ {
			fmt.Print(a[j], " ")
		}
		fmt.Println()
	}

}

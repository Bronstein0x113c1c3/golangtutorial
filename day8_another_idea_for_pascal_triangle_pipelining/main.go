package main

import "fmt"

func main() {
	// x := make([]int, 10)
	// x[0] = 1
	// result_chan := make(chan int)
	// index_chan := make(chan int)
	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		index_chan <- i
	// 	}
	// 	index_chan <- 0
	// 	close(index_chan)
	// }()
	// go func() {
	// 	for index := range index_chan {
	// 		if index == 0 {
	// 			close(result_chan)
	// 			break
	// 		} else {
	// 			result_chan <- x[index] + x[index-1]
	// 		}
	// 	}
	// }()

	// for result := range result_chan {
	// 	fmt.Print(result, " ")
	// }
	a := make([]uint64, 100)
	for i := 1; i <= 50; i++ {
		a = pascal_triangle_calculating(a)
		fmt.Println(a)
	}
	fmt.Println(a[25])
}

func pascal_triangle_calculating(a []uint64) []uint64 {
	x := make([]uint64, 1)
	x[0] = 1
	result_chan := make(chan uint64)
	index_chan := make(chan int)
	go func() {
		for i := 1; i <= len(a)-1; i++ {
			index_chan <- i
		}
		// index_chan <- 0
		close(index_chan)
	}()
	go func() {
		for index := range index_chan {
			// if index == 0 {
			// 	close(result_chan)
			// 	break
			// } else {
			result_chan <- a[index] + a[index-1]
			// }
		}
		close(result_chan)
	}()
	for result := range result_chan {
		x = append(x, result)
	}
	return x
}

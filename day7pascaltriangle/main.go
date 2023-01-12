package main

import (
	"fmt"
)

//	func pascal_trangle_calculating(in chan int) chan int {
//		c := make(chan int, 1)
//		var a, b int
//		go func() {
//			for {
//				a = <-in
//				b = <-in
//				c <- a + b
//			}
//		}()
//		return c
//	}
func main() {
	var b, c int
	a := make([]int, 13)
	a[0] = 1
	a[1] = 4
	a[2] = 6
	a[3] = 4
	a[4] = 1
	number_chan := make(chan int, 2)
	result_chan := make(chan int, 1000000)
	// time.Sleep(time.Millisecond * 0)
	// result_chan := make(chan int, 2)
	go func() {
		for {
			b = <-number_chan
			c = <-number_chan
			result_chan <- b + c
			// d = <-number_chan
			// result_chan <- c + d
			// a := <-pascal_trangle_calculating(number_chan)
			// fmt.Print(a)
		}
	}()
	for i := 1; i < 13; i++ {
		// number_chan <- a[i-2]
		number_chan <- a[i-1]
		number_chan <- a[i]
		if i >= 4 {
			a[i-3] = <-result_chan
		}

	}
	// fmt.Println(len(result_chan))
	// for i := 0; i < len(result_chan); i++ {
	// 	fmt.Println(<-result_chan)
	// }
	fmt.Println("result: ", a)
	fmt.Println(len(result_chan))
}
ko
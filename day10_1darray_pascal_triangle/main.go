package main

import "fmt"

func main() {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 0)
	for i := 1; i <= 50; i++ {
		prev_len := len(a)
		a = append(a, 1)
		for j := 1; j <= i; j++ {
			a = append(a, a[prev_len-(i+1)+j]+a[prev_len-(i+1)+j-1])
		}
		a = append(a, 0)
	}
	fmt.Println(a)
	for _, j := range a {
		if j == 0 {
			fmt.Println()
		} else {
			fmt.Printf("%v ", j)
		}
	}
}

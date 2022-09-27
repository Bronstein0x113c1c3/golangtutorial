package main

import "fmt"

func main() {
	points := [3]int{7, 10, 5}
	points[0] = 9
	points[1] = 10
	points[2] = 5
	fmt.Printf("%v, %T \n", points, points)
	fmt.Printf("%v, %T \n", points[2], points[2])

}

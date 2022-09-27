package main

import "fmt"

func main() {
	points := [3]string{"Alex", "Tom", "Yuh"}
	// points[0] =
	// points[1] =
	// points[2] =
	fmt.Printf("%v, %T \n", points, points)
	fmt.Printf("%v, %T \n", points[2], points[2])
	fmt.Println(len(points[1]))
}

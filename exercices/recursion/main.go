package main

import "fmt"

const m uint = 10e9 + 7

func Calculating(a, b uint) uint {

	if b == 0 {
		return 1

	} else {
		if b%2 == 0 {
			x := Calculating(a, b>>1) % m
			return (x * x) % m
		} else {
			x := (Calculating(a, b-1) % m)
			return (x * (a % m)) % m
		}
	}
}

func main() {
	var a, b uint
	fmt.Scan(&a, &b)
	fmt.Println(Calculating(a, b))
}

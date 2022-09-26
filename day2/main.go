package main

import "fmt"

func main() {
	var a int8 = 32
	var b int8 = 89
	fmt.Printf("%v \n", a+b)
	fmt.Printf("%v \n", a-b)
	fmt.Printf("%v \n", a/b)
	fmt.Printf("%v \n", a*b)
	fmt.Printf("%b \n", a^b)
	fmt.Printf("%b \n", a|b)
	fmt.Printf("%b \n", a&b)
	fmt.Printf("%b \n", a&^b)
}

/*
Primitive data types:
	1.Booleans
	2.Numerics
		-Integers
			signed integers: int8 int16 int32 int64
			unsigned integers: uint8 uint16 uint32 uint64
		-Float
		-Complex
	3.Text/File
		-String
		-Byte
		-Rune


*/

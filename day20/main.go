package main

import (
	"fmt"
	"reflect"
)

func main() {
	// s := "Hello"
	// s = "Bonjour!"
	// r := []rune(s)
	// r[0] = 54
	// s = string(r)
	s := "Â â"
	r1 := []rune(s)
	fmt.Println(reflect.TypeOf(r1[2]))
	fmt.Printf("%X \n", r1[2])
	fmt.Println("\U000000C2")
	// fmt.Println(utf8.DecodeRuneInString("Ă"))
}

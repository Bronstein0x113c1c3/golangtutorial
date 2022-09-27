package main

import "fmt"

func main() {
	month_map := map[string]int{
		"January":   31,
		"February":  28,
		"March":     31,
		"April":     30,
		"May":       31,
		"June":      30,
		"July":      31,
		"August":    31,
		"September": 30,
		"October":   31,
		"November":  30,
		"December":  31,
	}
	fmt.Println(month_map)             //the whole map
	fmt.Println(month_map["December"]) //access through key
	for k, v := range month_map {
		fmt.Printf("%v, %v \n", k, v)
	}
	//iterate the map through both key and value
	delete(month_map, "February")       //delete both key and value which key is "December"
	_, isExist := month_map["February"] //whether it exists or not?
	fmt.Println(isExist)
	month_map["February"] = 29 //re-append the February with 29 days
	fmt.Println(month_map)
}

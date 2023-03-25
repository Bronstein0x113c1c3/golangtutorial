package student

import (
	"fmt"
	"troy_canvas_optimized/command/entity"
)

func StudentCommand() {
	var s string
	Help()
	fmt.Println("Please give a command")
	fmt.Println("-----------------------")
	fmt.Scanln(&s)
	switch s {
	case "help":
		Help()
		StudentCommand()
	case "exit":
		return
	case "listinfo":
		fmt.Printf("Name: %v \n", entity.Student.Detail().Name)
		fmt.Printf("Canvas ID: %v \n", entity.Student.Detail().ID)
		fmt.Println()
		StudentCommand()
	}
}
func Help() {
	fmt.Println("Available command")
	fmt.Println("help")
	fmt.Println("exit")
	fmt.Println("listinfo")
}

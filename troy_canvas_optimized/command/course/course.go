package course

import (
	"fmt"
	"troy_canvas_optimized/command/entity"
)

func CourseCommand() {
	var s string
	Help()
	fmt.Println("Please give a command")
	fmt.Println("-----------------------")
	fmt.Scanln(&s)
	switch s {
	case "help":
		Help()
		CourseCommand()
	case "exit":
		return
	case "listall":
		fmt.Printf("%7v %30v\n", "ID", "Name")
		for _, c := range entity.Student.CoursesDetail() {
			fmt.Printf("%7v %30v\n", c.Detail().ID, c.Detail().Name)
		}
		CourseCommand()
	}

}
func Help() {
	fmt.Println("Available command")
	fmt.Println("help")
	fmt.Println("exit")
	fmt.Println("listall")
}

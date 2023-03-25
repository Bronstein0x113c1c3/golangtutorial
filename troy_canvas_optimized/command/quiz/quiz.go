package quiz

import (
	"fmt"
	"time"
	"troy_canvas_optimized/command/entity"
)

func QuizCommand() {
	var s string
	Help()
	fmt.Println("Please give a command")
	fmt.Println("-----------------------")
	fmt.Scanln(&s)
	switch s {
	case "help":
		Help()
		QuizCommand()
	case "exit":
		return
	case "listall":
		fmt.Printf("%30v   %30v    %30v \n", "ID", "Name", "Due")
		for _, c := range entity.Student.CoursesDetail() {
			for _, q := range c.QuizzesDetail() {
				fmt.Printf("%30v   %30v    %30v days ago   \n", q.Detail().ID, c.Detail().Name, int(-q.Detail().LockAt.Sub(time.Now()).Hours()/24))
			}
		}
		fmt.Println()
		QuizCommand()
	}
}
func Help() {
	fmt.Println("Available command")
	fmt.Println("help")
	fmt.Println("exit")
	fmt.Println("listall")
}

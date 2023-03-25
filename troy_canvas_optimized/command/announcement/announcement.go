package announcement

import (
	"fmt"
	"time"
	"troy_canvas_optimized/command/entity"

	"github.com/k3a/html2text"
)

func AnnouncementCommand() {
	var s string
	Help()
	fmt.Println("Please give a command")
	fmt.Println("-----------------------")
	fmt.Scanln(&s)

	switch s {
	case "help":
		Help()
		AnnouncementCommand()
	case "exit":
		return
	case "listall":
		fmt.Println()
		fmt.Printf("%7v  %15v  %20v \n", "ID", "Publisher", "Time")
		for _, c := range entity.Student.CoursesDetail() {
			for _, a := range c.AnnouncementsDetail() {
				fmt.Printf("%7v  %15v  %20v \n", a.Detail().ID, a.Detail().UserName, time.Now().Sub(a.Detail().PostedAt).Hours()/24)
			}
		}
		fmt.Println()
		AnnouncementCommand()
	case "content":
		for _, c := range entity.Student.CoursesDetail() {
			for _, a := range c.AnnouncementsDetail() {
				fmt.Println("******************")
				fmt.Printf("Course: %v \n", c.Detail().Name)
				fmt.Printf("Date %v \n", a.Detail().PostedAt.String())
				fmt.Printf("Content: \n %v \n", html2text.HTML2Text(a.Detail().Message))
				fmt.Println("******************")
			}
		}
		AnnouncementCommand()
	}

}
func Help() {
	fmt.Println("available command:")
	fmt.Println("help")
	fmt.Println("exit")
	fmt.Println("listall")
	fmt.Println("content")
}

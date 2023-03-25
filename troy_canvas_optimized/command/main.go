package command

import (
	"fmt"
	"os"
	"troy_canvas_optimized/command/announcement"
	"troy_canvas_optimized/command/course"
	"troy_canvas_optimized/command/quiz"
	"troy_canvas_optimized/command/student"
)

func Command() {
	help()
	fmt.Println("Please give a command")
	fmt.Println("-----------------------")
	var s string
	fmt.Scanln(&s)

	for {
		switch s {
		case "help":
			help()
			Command()
		case "student":
			fmt.Println("-----------------------")
			fmt.Println("Student Menu")
			student.StudentCommand()
			fmt.Println()
			Command()
		case "course":
			fmt.Println("-----------------------")
			fmt.Println("Course Menu")
			course.CourseCommand()
			fmt.Println()
			Command()

		case "quiz":
			fmt.Println("-----------------------")
			fmt.Println("Quiz Menu")
			quiz.QuizCommand()
			fmt.Println()
			Command()

		case "announcement":
			fmt.Println("-----------------------")
			fmt.Println("Announcement Menu")
			announcement.AnnouncementCommand()
			fmt.Println()
			Command()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("-----------------------")
			fmt.Println("not a command")
			fmt.Println()
			Command()
		}
	}
}

func help() {
	fmt.Println("Avalable command:")
	fmt.Println("student")
	fmt.Println("course")
	fmt.Println("announcement")
	fmt.Println("quiz")
	fmt.Println("exit")

}

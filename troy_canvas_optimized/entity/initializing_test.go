package entity_test

import (
	"fmt"
	"os"
	"testing"
	"troy_canvas_optimized/entity"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	_ = godotenv.Load("../env/.env")
	token := os.Getenv("TOKEN")
	fmt.Println(token)
	Student := entity.Init(token)
	//test the student info
	if Student.Detail().ID == 0 {
		t.Error("Could not get student detail")
		return
	} else {
		fmt.Println(Student.Detail())
	}
	fmt.Println()
	fmt.Println("Passed the student info test")
	fmt.Println()

	//test courses
	for _, c := range Student.CoursesDetail() {
		fmt.Println(c.Detail())
	}
	fmt.Println()
	fmt.Println("Passed courses info test")
	fmt.Println()
	//test quizzes
	for _, c := range Student.CoursesDetail() {
		for _, q := range c.QuizzesDetail() {
			fmt.Println(q.Detail())
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("Passes quizzes test")
	fmt.Println()

	//test announements
	for _, c := range Student.CoursesDetail() {
		for _, a := range c.AnnouncementsDetail() {
			fmt.Println(a.Detail())
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("Passes announcements test")
	fmt.Println()

}

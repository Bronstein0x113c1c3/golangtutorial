package implementation

import (
	_ "fmt"
	_ "net/http"
	"troy_canvas/entity"
)

type StudentInfo struct {
	Info    *entity.Student_Info_Instructure
	Courses []*Course
}

type Course struct {
	Info          *entity.Course_Info_Instructure
	Quizzes       []*Quiz
	Announcements []*Announcement
}
type Quiz struct {
	Info *entity.Quiz_Instructure
}
type Announcement struct {
	Info *entity.Course_Announement_Instructure
}

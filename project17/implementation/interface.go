package implementation

import (
	"project17/entity"
)

type StudentInterface interface {
	AddStudents(list *entity.StudentList, student_info *entity.Student) *entity.StudentList
	GetStudents() *entity.StudentList
	DeleteStudents(list *entity.StudentList, id string, student_info *entity.Student) *entity.StudentList
	UpdateStudents(list *entity.StudentList, id string, student_info *entity.Student) *entity.StudentList
}

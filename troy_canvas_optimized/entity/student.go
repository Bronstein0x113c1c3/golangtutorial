package entity

import (
	"fmt"
	"troy_canvas_optimized/common"
	"troy_canvas_optimized/entity/instructure"
)

type Student struct {
	info    instructure.Student_Info_Instructure
	courses []*Course
}

func (s *Student) importing(info instructure.Student_Info_Instructure) {
	s.info = info
}
func (s *Student) getallcourses(token string) {

	//First, initiate courses list first
	courses := []instructure.Course_Info_Instructure{}
	//Get the value of the courses list
	common.Get(common.CourseInfoEndpoint, token, &courses)
	fmt.Println(courses)
	//for each course in courses list, create a course class, import the course data then appending courses field in courses
	for _, course := range courses {
		c := &Course{}
		c.importing(course)
		s.courses = append(s.courses, c)
	}
}
func (s *Student) Detail() instructure.Student_Info_Instructure {
	return s.info
}
func (s *Student) CoursesDetail() []*Course {
	return s.courses
}

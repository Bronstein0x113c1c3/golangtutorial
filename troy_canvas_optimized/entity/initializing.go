package entity

import (
	"troy_canvas_optimized/common"
	"troy_canvas_optimized/entity/instructure"
)

func Init(token string) *Student {
	student := &Student{}
	info := &instructure.Student_Info_Instructure{}
	common.Get(common.StudentInfoEndpoint, token, info)
	student.importing(*info)
	student.getallcourses(token)
	for _, c := range student.courses {
		c.getallannouncements(token)
		c.getallquizzes(token)
	}
	return student
}

package entity

import (
	"fmt"
	"troy_canvas_optimized/common"
	"troy_canvas_optimized/entity/instructure"
)

type Course struct {
	info          instructure.Course_Info_Instructure
	quizzes       []*Quiz
	announcements []*Announcement
}

func (c *Course) importing(info instructure.Course_Info_Instructure) {
	c.info = info
}
func (c *Course) getallannouncements(token string) {
	//First, initiate courses list first
	announcements := []instructure.Announcement_Info_Instructure{}
	//Get the value of the courses list
	common.Get(fmt.Sprintf(common.AnnouncementInfoEndpoint, c.info.ID), token, &announcements)
	//for each course in courses list, create a course class, import the course data then appending courses field in courses
	for _, announcement := range announcements {
		a := &Announcement{}
		a.importing(announcement)
		c.announcements = append(c.announcements, a)
	}
}
func (c *Course) getallquizzes(token string) {
	var quizzes []instructure.Quiz_Info_Instructure
	common.Get(fmt.Sprintf(common.QuizInfoEndpoint, c.info.ID), token, &quizzes)
	for _, quiz := range quizzes {
		q := &Quiz{}
		q.importing(quiz)
		c.quizzes = append(c.quizzes, q)
	}
}
func (c *Course) Detail() instructure.Course_Info_Instructure {
	return c.info
}
func (c *Course) QuizzesDetail() []*Quiz {
	return c.quizzes
}
func (c *Course) AnnouncementsDetail() []*Announcement {
	return c.announcements
}

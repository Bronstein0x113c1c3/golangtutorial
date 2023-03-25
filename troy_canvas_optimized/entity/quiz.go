package entity

import "troy_canvas_optimized/entity/instructure"

type Quiz struct {
	info instructure.Quiz_Info_Instructure
}

func (q *Quiz) importing(info instructure.Quiz_Info_Instructure) {
	q.info = info
}
func (q *Quiz) Detail() instructure.Quiz_Info_Instructure {
	return q.info
}

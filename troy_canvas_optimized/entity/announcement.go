package entity

import "troy_canvas_optimized/entity/instructure"

type Announcement struct {
	info instructure.Announcement_Info_Instructure
}

func (a *Announcement) importing(info instructure.Announcement_Info_Instructure) {
	a.info = info
}
func (a *Announcement) Detail() instructure.Announcement_Info_Instructure {
	return a.info
}
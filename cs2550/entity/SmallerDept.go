package entity

type SmallerDept struct {
	Name         string
	Manager      Human
	Vice_Manager Human
	Staffs       []*Human
}

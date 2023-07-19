package entity

type Company struct {
	President   Human
	Name        string
	Departments []*Department
}

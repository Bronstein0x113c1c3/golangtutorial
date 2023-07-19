package entity

type Department struct {
	Name         string
	CEO          Human
	Vice_CEO     Human
	SmallerDepts []*SmallerDept
}

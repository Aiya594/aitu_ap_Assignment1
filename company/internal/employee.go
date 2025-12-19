package internal

const (
	PartTime = "PartTimeEmployee"
	FullTime = "FullTimeEmployee"
)

type IEmployee interface {
	GetDetail() EmployeeDetails
}

type EmployeeDetails struct {
	ID   uint64
	Name string
	Type string
	Pay  float64
}

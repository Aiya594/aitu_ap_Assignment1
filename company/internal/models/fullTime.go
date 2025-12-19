package models

import "github.com/Aiya594/aitu_ap_Assignment1/company/internal"

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Role   string
	Salary float64
}

func NewFullTimeEmployee(id uint64, name, role string, salary float64) *FullTimeEmployee {
	return &FullTimeEmployee{
		ID:     id,
		Name:   name,
		Role:   role,
		Salary: salary,
	}
}

func (f *FullTimeEmployee) GetDetail() internal.EmployeeDetails {
	return internal.EmployeeDetails{
		ID:   f.ID,
		Name: f.Name,
		Type: internal.FullTime,
		Pay:  f.Salary,
	}
}

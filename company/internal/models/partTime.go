package models

import "github.com/Aiya594/aitu_ap_Assignment1/company/internal"

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	Role        string
	HourPayment float64
}

func NewPartTimeEmployee(id uint64, name, role string, hourPayment float64) *PartTimeEmployee {
	return &PartTimeEmployee{
		ID:          id,
		Name:        name,
		Role:        role,
		HourPayment: hourPayment,
	}
}

func (p *PartTimeEmployee) GetDetail() internal.EmployeeDetails {
	return internal.EmployeeDetails{
		ID:   p.ID,
		Name: p.Name,
		Type: internal.PartTime,
		Pay:  p.HourPayment,
	}
}

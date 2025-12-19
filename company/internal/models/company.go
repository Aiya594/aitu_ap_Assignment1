package models

import (
	"fmt"

	"github.com/Aiya594/aitu_ap_Assignment1/company/internal"
)

type Company struct {
	Name      string
	Employees map[uint64]internal.IEmployee
}

func NewCompany(name string) *Company {
	return &Company{
		Name:      name,
		Employees: make(map[uint64]internal.IEmployee)}
}

func (c *Company) AddEmployee(e internal.IEmployee) {
	newID := uint64(len(c.Employees) + 1)
	c.Employees[newID] = e
}

func (c *Company) ListEmployees() {
	for _, em := range c.Employees {
		fmt.Println(em.GetDetail())
	}
}

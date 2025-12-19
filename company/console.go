package company

import (
	"fmt"

	"github.com/Aiya594/aitu_ap_Assignment1/company/internal"
	"github.com/Aiya594/aitu_ap_Assignment1/company/internal/models"
)

func RunCompany() {
	company := models.NewCompany("MyCompany")

	for {
		fmt.Println(`
This is Employee Management System
Options:
	1.List all employees of company
	2.Show details of cocrete employee
	3.Add employee to company 
	4.Stop`)

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("couldnt read the input:", err)
			continue
		}

		switch choice {
		case 1:
			company.ListEmployees()

		case 2:
			fmt.Println("Please enter the id of employee")
			var id uint64
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			em, exists := company.Employees[id]
			if !exists {
				fmt.Println("No such employee in company")
				continue
			}
			fmt.Println(em.GetDetail())

		case 3:
			fmt.Println("Choose the employee type(f - full time, p - part time):")
			var emType string
			_, err := fmt.Scan(&emType)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}

			fmt.Println("To add new employee to company enter these values(name, role, salary(or hour payment))")
			var name, role string
			var salary float64
			_, err = fmt.Scan(&name, &role, &salary)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			id := uint64(len(company.Employees) + 1)
			var newEm internal.IEmployee

			if emType == "f" {
				newEm = models.NewFullTimeEmployee(id, name, role, salary)
			} else if emType == "p" {
				newEm = models.NewPartTimeEmployee(id, name, role, salary)
			} else {
				fmt.Println("Invalid employee type")
				continue
			}

			company.AddEmployee(newEm)
			fmt.Println("Added new employee to company")

		case 4:
			return

		default:
			fmt.Println("Invalid choice")
			continue
		}

	}
}

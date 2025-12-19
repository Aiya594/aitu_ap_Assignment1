package main

import (
	"fmt"

	"github.com/Aiya594/aitu_ap_Assignment1/bank"
	"github.com/Aiya594/aitu_ap_Assignment1/company"
	"github.com/Aiya594/aitu_ap_Assignment1/library"
	"github.com/Aiya594/aitu_ap_Assignment1/shapes"
)

func main() {
	fmt.Println(`
Assignment1:
	1.Run Library Management System console
	2.Run Shapes & Interfaces console
	3.Run Employee Management System console 
	4.Run Bank Account Simulation console`)
	var choice int
	fmt.Println("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("couldnt read the input:", err)
	}
	switch choice {
	case 1:
		library.RunLibrary()
	case 2:
		shapes.RunShapes()
	case 3:
		company.RunCompany()
	case 4:
		bank.RunBank()
	default:
		fmt.Println("Unrecognized choice")
	}

}

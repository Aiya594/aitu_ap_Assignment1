package bank

import (
	"fmt"

	"github.com/Aiya594/aitu_ap_Assignment1/bank/internal/models"
)

func RunBank() {

	ba := models.NewBankAccount("user1", "Aiya", 10000)

	for {
		fmt.Println(`
This is Bank Account Simulation 
	1. Deposit
	2. Withdraw
	3. Show balance
	4. Show transactions
	5. Exit`)

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("couldnt read the input:", err)
			continue
		}
		switch choice {
		case 1:
			fmt.Println("Enter deposit amount: ")
			var amount float64
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			message, ok := ba.Deposit(amount)
			if !ok {
				fmt.Println(message)
				continue
			}
			fmt.Println("Successfully deposit: ", amount)
		case 2:
			fmt.Println("Enter withdraw amount: ")
			var amount float64
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println("couldnt read the input:", err)
				continue
			}
			message, ok := ba.Withdraw(amount)
			if !ok {
				fmt.Println(message)
				continue
			}
			fmt.Println("Successfully withdraw: ", amount)

		case 3:
			fmt.Println("Your balance is: ", ba.GetBalance())

		case 4:
			fmt.Println("History of transactions:")
			transactions := ba.TransactionsHistory
			for _, transaction := range transactions {
				fmt.Println(transaction)
			}

		case 5:
			fmt.Println("Bye")
			return
			
		default:
			fmt.Println("Unrecognized choice")
			continue
		}

	}
}

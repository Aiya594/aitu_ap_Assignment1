package models

type BankAccount struct {
	ID                  string
	UserName            string
	Balance             float64
	TransactionsHistory []Transaction
}

func NewBankAccount(id string, userName string, balance float64) *BankAccount {
	return &BankAccount{
		ID:       id,
		UserName: userName,
		Balance:  balance,
	}
}

func (ba *BankAccount) Deposit(amount float64) (string, bool) {
	if amount <= 0 {
		return "amount must be greater that 0", false
	}
	transaction := Transaction{
		Type:   deposit,
		Amount: amount,
	}
	ba.TransactionsHistory = append(ba.TransactionsHistory, transaction)

	ba.Balance += amount
	return "transaction saved in history", true
}

func (ba *BankAccount) Withdraw(amount float64) (string, bool) {
	if amount <= 0 {
		return "amount must be greater that 0", false
	} else if amount > ba.Balance {
		return "not enough balance", false
	}
	transaction := Transaction{
		Type:   withdraw,
		Amount: amount,
	}
	ba.TransactionsHistory = append(ba.TransactionsHistory, transaction)
	ba.Balance -= amount
	return "transaction saved in history", true
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.Balance
}

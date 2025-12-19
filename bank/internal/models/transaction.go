package models

type Transaction struct {
	Type   string
	Amount float64
}

// transacttion types
const (
	deposit  = "Deposit"
	withdraw = "Withdraw"
)

package domain

import "time"

type Transaction struct {
	TransactionId string    `bson:"transaction_id" json:"transaction_id,omitempty"`
	FromAccountId string    `bson:"from_account_id" json:"from_account_id,omitempty"`
	ToAccountId   string    `bson:"to_account_id" json:"to_account_id,omitempty"`
	Action        string    `bson:"action" json:"action,omitempty"`
	Amount        int       `bson:"amount" json:"amount,omitempty"`
	BalanceChange int       `bson:"balance_change" json:"balance_change,omitempty"`
	DateIssued    time.Time `bson:"date_issued" json:"date_issued,omitempty"`
	CreatedAt     string    `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt     string    `bson:"updated_at" json:"updated_at,omitempty"`
}
type Action int

const (
	Deposit  Action = 0
	Withdraw Action = 1
	Transfer Action = 2
)

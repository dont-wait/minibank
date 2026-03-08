package domain

import "time"

type Account struct {
	AccountId  string    `bson:"account_id" json:"account_id,omitempty"`
	Balance    int       `bson:"balance" json:"balance,omitempty"`
	PinCode    string    `bson:"pin_code" json:"pin_code,omitempty"`
	Status     Status    `bson:"status" json:"status,omitempty"`
	Type       Type      `bson:"type" json:"type,omitempty"`
	CustomerId string    `bson:"customer_id" json:"customer_id,omitempty"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
type Status int

const (
	InActive Status = 0
	Active   Status = 1
	Frozen   Status = -1
	Closed   Status = -2
)

type Type int

const (
	Checking Type = 0
	Savings  Type = 1
)
